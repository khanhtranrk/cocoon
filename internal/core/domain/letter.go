package domain

type LetterType uint8
type LetterStatus uint8

const (
  Request  LetterType = 1
  Response LetterType = 2

  SystemError  LetterStatus = 11
  LetterError  LetterStatus = 12

  Waiting      LetterStatus = 21
  Keeping      LetterStatus = 22
  Pending      LetterStatus = 23
  Done         LetterStatus = 24
  MessageError LetterStatus = 25

  Sent         LetterStatus = 31
)

type Letter struct {
  Id         uint64
  Type       LetterType
  Code       uint64
  ForeignId  uint64
  CommitTime uint64
  Message    []byte
  Status     LetterStatus
}
