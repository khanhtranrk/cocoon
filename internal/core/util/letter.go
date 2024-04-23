package util

import (
	"encoding/binary"
	"fmt"

	"github.com/khanhtranrk/cocoon/internal/core/domain"
)

func LetterToDeliverable(lt *domain.Letter, selfId uint64, typ uint8) ([]byte, error) {
  var typEnc uint8 = typ

  codeEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(codeEnc, lt.Code)

  senderIdEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(senderIdEnc, selfId)

  receiverIdEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(receiverIdEnc, lt.ForeignId)

  commitTimeEnc := make([]byte, 8)
  binary.BigEndian.PutUint64(commitTimeEnc, lt.CommitTime)

  lenOfMsgEnc := make([]byte, 4)
  binary.BigEndian.PutUint32(lenOfMsgEnc, uint32(len(lt.Message)))

  var dv []byte
  dv = append(dv, typEnc)
  dv = append(dv, codeEnc...)
  dv = append(dv, senderIdEnc...)
  dv = append(dv, receiverIdEnc...)
  dv = append(dv, commitTimeEnc...)
  dv = append(dv, lenOfMsgEnc...)
  dv = append(dv, lt.Message...)

  return dv, nil
}

func DeliverableToLetter(dv []byte) (uint8, *domain.Letter, error) {
  if len(dv) < 37 {
    return 0, nil, fmt.Errorf("Len of Deliverable was wrong!")
  }

  typ := dv[0]
  code := binary.BigEndian.Uint64(dv[1:9])
  foreignId := binary.BigEndian.Uint64(dv[9:17])
  commitTime := binary.BigEndian.Uint64(dv[25:33])
  lenOfMsg := binary.BigEndian.Uint32(dv[33:37])

  if len(dv) < 37 + int(lenOfMsg) {
    return 0, nil, fmt.Errorf("Len of Deliverable was wrong!")
  }

  msg := dv[37:37 + lenOfMsg]

  return typ, &domain.Letter{
    Code: code,
    ForeignId: foreignId,
    CommitTime: commitTime,
    Message: msg,
  }, nil
}
