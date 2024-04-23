package port

import "github.com/khanhtranrk/cocoon/internal/core/domain"

type CitizenRepository interface {
  ListAllCitizens() ([]*domain.Citizen, error)
  GetCitizenById(id uint64) (*domain.Citizen, error)
}

type CitizenService interface {
  ListAllCitizens() ([]*domain.Citizen, error)
  GetCitizenById(id uint64) (*domain.Citizen, error)
}
