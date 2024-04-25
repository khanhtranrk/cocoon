package repository

import (
	"database/sql"
	"fmt"

	"github.com/khanhtranrk/cocoon/external/core/domain"
)

type ContactRepository struct {
  db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
  return &ContactRepository{db}
}

func (cr *ContactRepository) ListAllContacts() ([]*domain.Contact, error) {
  query := "SELECT citizen_id, permission FROM contacts"
  rows, err := cr.db.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var contacts []*domain.Contact
  for rows.Next() {
    var contact domain.Contact

    err = rows.Scan(&contact.CitizenId, &contact.Gate)
    if err != nil {
      return nil, err
    }

    contacts = append(contacts, &contact)
  }

  return contacts, nil
}

func (cr *ContactRepository) GetContactByCitizenId(id uint64) (*domain.Contact, error) {
  query := "SELECT citizen_id, gate FROM contacts WHERE citizen_id = ?"
  rows, err := cr.db.Query(query, id)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  if !rows.Next() {
    return nil, fmt.Errorf("Contact not found: %v", id)
  }

  var contact domain.Contact
  err = rows.Scan(&contact.CitizenId, &contact.Gate)
  if err != nil {
    return nil, err
  }

  return &contact, nil
}
