package repository

import (
	"database/sql"
	"fmt"

	"github.com/khanhtranrk/cocoon/external/core/domain"
)

type CitizenRepository struct {
  db *sql.DB
}

func NewCitizenRepository(db *sql.DB) *CitizenRepository {
  return &CitizenRepository{db}
}

func (cr *CitizenRepository) ListAllCitizens() ([]*domain.Citizen, error) {
  query := `SELECT id, name, contact_gate, legal_gate, registration_date FROM citizens`
  rows, err := cr.db.Query(query)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  var citizens []*domain.Citizen
  for rows.Next() {
    var citizen domain.Citizen

    err = rows.Scan(
      &citizen.Id,
      &citizen.Name,
      &citizen.ContactGate,
      &citizen.LegalGate,
      &citizen.RegistrationDate,
    )

    if err != nil {
      return nil, err
    }

    citizens = append(citizens, &citizen)
  }

  return citizens, nil
}

func (cr *CitizenRepository) GetCitizenById(id uint64) (*domain.Citizen, error) {
  query := `SELECT id, name, contact_gate, legal_gate, registration_date
            FROM citizens
            WHERE id = ?`

  rows, err := cr.db.Query(query, id)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  if !rows.Next() {
    return nil, fmt.Errorf("Citizen not found: %v", id)
  }

  var citizen domain.Citizen
  err = rows.Scan(
    &citizen.Id,
    &citizen.Name,
    &citizen.ContactGate,
    &citizen.LegalGate,
    &citizen.RegistrationDate,
  )

  if err != nil {
    return nil, err
  }

  return &citizen, nil
}
