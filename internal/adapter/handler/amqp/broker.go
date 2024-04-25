package amqp

import (
	"log"

	"github.com/khanhtranrk/cocoon/internal/adapter/config"
	"github.com/streadway/amqp"
)

type Broker struct {
  Connection *amqp.Connection
  Channel *amqp.Channel
}

func New(cfg *config.Config) (*Broker, error) {
  conn, err := amqp.Dial(cfg.BrokerUrl)
  if err != nil {
    return nil, err
  }

  channel, err := conn.Channel()
  if err != nil {
    return nil, err
  }

  return &Broker{Connection: conn, Channel: channel}, nil
}

func (b *Broker) Listen(queue string, outChan chan []byte) {
  msgs, err := b.Channel.Consume(queue, "", true, false, false, false, nil)
  if err != nil {
    log.Fatalf(err.Error())
  }

  for d := range msgs {
    outChan <-d.Body
  }
}

func (b *Broker) SendMessage(queue string, msg []byte) error {
  return b.Channel.Publish("", queue, false, false, amqp.Publishing{
      ContentType: "text/plain",
      Body: msg,
  })
}

// HACK: cannot call 2 times if chanel.Close() error
func (b *Broker) Close() error {
  err := b.Channel.Close()

  if err != nil {
    return err
  }

  err = b.Connection.Close()

  if err != nil {
    return err
  }

  return nil
}
