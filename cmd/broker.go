package main

import (
	"log"
	"github.com/streadway/amqp"
)

type Broker struct {
  hub *Hub
}

func NewBroker(hub *Hub) *Broker {
  return &Broker{hub: hub}
}

func (b *Broker) Listen() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
    if err != nil {
        log.Fatalf(err.Error())
    }
    defer conn.Close()

    channel, err := conn.Channel()
    if err != nil {
        log.Fatalf(err.Error())
    }
    defer channel.Close()

    msgs, err := channel.Consume("mavis", "rabit", true, false, false, false, nil)
    if err != nil {
        log.Fatalf(err.Error())
    }

    for d := range msgs {
      b.hub.broadcast <-d.Body
    }
}

func (b *Broker) SendMessage(message []byte) {
  conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
  if err != nil {
      log.Fatalf(err.Error())
  }
  defer conn.Close()

  channel, err := conn.Channel()
  if err != nil {
      log.Fatalf(err.Error())
  }
  defer channel.Close()

  err = channel.Publish("", "test", false, false, amqp.Publishing{
      ContentType: "text/plain",
      Body: message,
  })
}
