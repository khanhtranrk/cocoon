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

type CocoonHandlerFuncs interface {
  MessageHandler(msg []byte) (uint8, []byte, error)
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
    SendLetterChan: make(chan *domain.Letter),
    TerminateChan: make(chan *bool),
  }, nil
}

// HACK: queue "taistra" is fixed value
// OPTIMIZE: This code snippet violates the DRY principle
func (c *Cocoon) SendRequestLetter(lt *Letter) error {
  var sigEnc uint8 = 24
  var verEnc uint8 = 1
  var typEnc uint8 = 1

  codeEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(codeEnc, lt.Code)

  senderIdEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(senderIdEnc, c.Cert.Id)

  receiverIdEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(receiverIdEnc, lt.ForeignId)

  commitTimeEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(commitTimeEnc, lt.CommitTime)

  lenOfMsgEnc := make([]byte, 4)
  binary.BigEndian.PutUint32(lenOfMsgEnc, uint32(len(lt.Message)))

  var dv []byte
  dv = append(dv, sigEnc)
  dv = append(dv, verEnc)
  dv = append(dv, typEnc)
  dv = append(dv, codeEnc...)
  dv = append(dv, senderIdEnc...)
  dv = append(dv, receiverIdEnc...)
  dv = append(dv, commitTimeEnc...)
  dv = append(dv, lenOfMsgEnc...)
  dv = append(dv, lt.Message...)

  // find contact to give gate
  ct, err := c.ContactService.GetContactByCitizenId(lt.ForeignId)
  if err != nil {
    return err
  }

  if ct != nil {
    return fmt.Errorf("SendResponseLetter: Could not find contact with CitizenId = %v", lt.ForeignId)
  }

  return c.Broker.SendMessage(ct.Gate, dv)
}

// HACK: queue "taistra" is fixed value
// OPTIMIZE: This code snippet violates the DRY principle
func (c *Cocoon) SendResponseLetter(lt *Letter) error {
  var sig uint8 = 24
  var ver uint8 = 1
  lt.Type = 2

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

  // find contact to give gate
  ct, err := c.ContactService.GetContactByCitizenId(lt.ForeignId)
  if err != nil {
    return err
  }

  if ct != nil {
    return fmt.Errorf("SendResponseLetter: Could not find contact with CitizenId = %v", lt.ForeignId)
  }

  return c.Broker.SendMessage(ct.Gate, dv)
}
