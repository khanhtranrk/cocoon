package domain

type Letter struct {
  Id uint64
  Code uint64
  ForeignId uint64
  CommitTime uint64
  Message []byte
  Status uint8
}
