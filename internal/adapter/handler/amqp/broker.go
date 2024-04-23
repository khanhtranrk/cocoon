package amqp

import (
	"github.com/khanhtranrk/cocoon/internal/adapter/config"
	"github.com/streadway/amqp"
)

type Broker struct {
  conn *amqp.Connection
  channel *amqp.Channel
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

    return &Broker{conn: conn, channel: channel}, nil
}

func (b *Broker) Close() error {
  err := b.channel.Close()

  if err != nil {
    return err
  }

  err = b.conn.Close()

  if err != nil {
    return err
  }

  return nil
}

