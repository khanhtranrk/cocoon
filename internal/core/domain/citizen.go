package domain

import "time"

type Citizen struct {
  Id uint64
  Name string
  ContactGate string
  RegistrationDate time.Time
}
