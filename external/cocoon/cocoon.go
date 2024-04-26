package cocoon

import (
	"database/sql"
	"encoding/binary"
	"fmt"

	"github.com/khanhtranrk/cocoon/internal/adapter/certificate"
	"github.com/khanhtranrk/cocoon/internal/adapter/config"
	"github.com/khanhtranrk/cocoon/internal/adapter/handler/amqp"
	"github.com/khanhtranrk/cocoon/internal/adapter/storage/sqlite"
	"github.com/khanhtranrk/cocoon/internal/core/domain"
	"github.com/khanhtranrk/cocoon/internal/core/port"
)

type Citizen = domain.Citizen
type Contact = domain.Contact
type Letter = domain.Letter
type LetterType = domain.LetterType
type LetterStatus = domain.LetterStatus

const (
  RequestType  LetterType = domain.Request
  ResponseType LetterType = domain.Response

  SystemErrorStatus  LetterStatus = domain.SystemError
  LetterErrorStatus  LetterStatus = domain.LetterError

  WaitingStatus      LetterStatus = domain.Waiting
  KeepingStatus      LetterStatus = domain.Keeping
  PendingStatus      LetterStatus = domain.Pending
  DoneStatus         LetterStatus = domain.Done
  MessageErrorStatus LetterStatus = domain.MessageError

  SentStatus         LetterStatus = domain.Sent
)

type CocoonHandlerFuncs interface {
  MessageHandler(msg []byte) (LetterStatus, []byte, error)
}

type Cocoon struct {
  // Base
  Config *config.Config
  Db *sql.DB
  Broker *amqp.Broker
  Cert *certificate.Certificate

  // Service
  ContactService port.ContactService
  LetterService port.LetterService

  // Chanels
  SuspiciousLetterChan chan *Letter
  ProcessRequestLetterChan chan *Letter
  ProcessResponseLetterChan chan *Letter
  SendRequestLetterChan chan *Letter
  SendResponseLetterChan chan *Letter
  TerminateChan chan *bool

  // Funcs
  CocoonHandlerFuncs
}

func New() (*Cocoon, error) {
  // load certificate
  cert, err := certificate.LoadAbilityCertificate()
  if err != nil {
    return nil, err
  }

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
    Cert: cert,
    SuspiciousLetterChan: make(chan *domain.Letter),
    ProcessRequestLetterChan: make(chan *domain.Letter),
    ProcessResponseLetterChan: make(chan *domain.Letter),
    SendRequestLetterChan: make(chan *domain.Letter),
    SendResponseLetterChan: make(chan *domain.Letter),
    TerminateChan: make(chan *bool),
  }, nil
}

func (c *Cocoon) SendLetter(lt *Letter) error {
  // header
  var sig uint8 = 24
  var ver uint8 = 1

  // Delivery infos
  code := make([]byte, 8)
  binary.BigEndian.PutUint64(code, lt.Code)

  senderId := make([]byte, 8)
  binary.BigEndian.PutUint64(senderId, c.Cert.Id)

  receiverId := make([]byte, 8)
  binary.BigEndian.PutUint64(receiverId, lt.ForeignId)

  commitTime := make([]byte, 8)
  binary.BigEndian.PutUint64(commitTime, lt.CommitTime)

  lenOfMsg := make([]byte, 4)
  binary.BigEndian.PutUint32(lenOfMsg, uint32(len(lt.Message)))

  // letter
  var dv []byte
  dv = append(dv, sig)
  dv = append(dv, ver)
  dv = append(dv, lt.Type)
  dv = append(dv, code...)
  dv = append(dv, senderId...)
  dv = append(dv, receiverId...)
  dv = append(dv, commitTime...)
  dv = append(dv, lenOfMsg...)
  dv = append(dv, lt.Message...)

  // find gate of foreign
  ct, err := c.ContactService.GetContactByCitizenId(lt.ForeignId)
  if err != nil {
    return err
  }

  if ct != nil {
    return fmt.Errorf("SendLetter: Could not find contact with CitizenId = %v", lt.ForeignId)
  }

  return c.Broker.SendMessage(ct.Gate, dv)
}
