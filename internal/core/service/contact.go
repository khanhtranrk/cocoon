package service

import (
	"github.com/khanhtranrk/cocoon/internal/core/domain"
	"github.com/khanhtranrk/cocoon/internal/core/port"
)

type ContactService struct {
  rp port.ContactRepository
}

func NewContactService(rp port.ContactRepository) *ContactService {
  return &ContactService{rp}
}

func (cs *ContactService) ListAllContacts() ([]*domain.Contact, error) {
  return cs.rp.ListAllContacts()
}

func (cs *ContactService) GetContactByCitizenId(contactId uint64) (*domain.Contact, error) {
  return cs.rp.GetContactByCitizenId(contactId)
}
