package port

import "github.com/khanhtranrk/cocoon/internal/core/domain"

type ContactRepository interface {
  ListAllContacts() ([]*domain.Contact, error)
  GetContactByCitizenId(id uint64) (*domain.Contact, error)
}

type ContactService interface {
  ListAllContacts() ([]*domain.Contact, error)
  GetContactByCitizenId(id uint64) (*domain.Contact, error)
}
