package cocoon

import (
	"encoding/binary"
	"log"
)


func (c *Cocoon) Keeper(dv []byte) {
  // check minimal length
  if len(dv) < 39 {
    log.Printf("Letter's length is invalid. It is discarded")
    return
  }

  // extract
  sig := dv[0]
  ver := dv[1]
  typ := dv[2]
  code := binary.BigEndian.Uint64(dv[3:11])
  senderId := binary.BigEndian.Uint64(dv[11:19])
  receiverId := binary.BigEndian.Uint64(dv[19:27])
  commitTime := binary.BigEndian.Uint64(dv[27:35])
  lenOfMsg := binary.BigEndian.Uint32(dv[35:39])

  if len(dv) != 39 + int(lenOfMsg) {
    log.Printf("Letter's length is invalid. It is discarded")
    return
  }

  msg := dv[39:39 + lenOfMsg]

  lt := &Letter{
    Type: typ,
    Code: code,
    ForeignId: senderId,
    CommitTime: commitTime,
    Message: msg,
  }

  // Checking
  if sig != 24 || ver != 1 || !(typ == 1 || typ == 2) {
    c.SuspiciousLetterChan <-lt
  }

  ct, err := c.ContactService.GetContactByCitizenId(senderId)
  if err != nil {
    c.SuspiciousLetterChan <-lt
  }

  if ct == nil {
    c.SuspiciousLetterChan <-lt
  }

  if receiverId != c.Cert.Id {
    c.SuspiciousLetterChan <-lt
  }

  if typ == 1 {
    c.ProcessRequestLetterChan <-lt
  } else {
    c.ProcessResponseLetterChan <-lt
  }
}

func (c *Cocoon) ListenContactGate()  {
  msgs, err := c.Broker.Channel.Consume(c.Cert.ContactGate, "", true, false, false, false, nil)
  if err != nil {
      log.Fatalf(err.Error())
  }

  for d := range msgs {
    c.Keeper(d.Body)
  }
}
