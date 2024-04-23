package cocoon

import (
	"database/sql"

	"github.com/khanhtranrk/cocoon/internal/adapter/config"
	"github.com/khanhtranrk/cocoon/internal/adapter/handler/amqp"
	"github.com/khanhtranrk/cocoon/internal/adapter/storage/sqlite"
	"github.com/khanhtranrk/cocoon/internal/core/domain"
	"github.com/khanhtranrk/cocoon/internal/core/port"
)

type Citizen = domain.Citizen
type Contact = domain.Contact
type Letter = domain.Letter

type CocoonHandleFuncs interface {
}

type Cocoon struct {
  // Base
  Config *config.Config
  Db *sql.DB
  Broker *amqp.Broker

  // Service
  ContactService *port.ContactService
  LetterService *port.LetterService

  // Chanels
  UnprocessedLetterChan chan *Letter
  UnsentLetterChan chan *Letter
  ProcessLetterChan chan *Letter
  SendLetterChan chan *Letter
  TerminateChan chan *bool

  // Funcs
  CocoonHandleFuncs
}

func New() (*Cocoon, error) {
  // load config
  cfg, err := config.New()
  if err != nil {
    return nil, err
  }

  // connect to data
  db, err := sqlite.New(cfg)
  if err != nil {
    return nil, err
  }

  // connect to broker
  broker, err := amqp.New(cfg)
  if err != nil {
    return nil, err
  }

  return &Cocoon{
    Config: cfg,
    Db: db,
    Broker: broker,
    UnprocessedLetterChan: make(chan *domain.Letter),
    UnsentLetterChan: make(chan *domain.Letter),
    ProcessLetterChan: make(chan *domain.Letter),
    SendLetterChan: make(chan *domain.Letter),
    TerminateChan: make(chan *bool),
  }, nil
}

func (cc *Cocoon) SendLetter(lt *Letter) error {
}
