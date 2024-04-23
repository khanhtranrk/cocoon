package service

import (
	"github.com/khanhtranrk/cocoon/internal/core/domain"
	"github.com/khanhtranrk/cocoon/internal/core/port"
)

type CitizenService struct {
  rp port.CitizenRepository
}

func NewCitizenService(rp port.CitizenRepository) *CitizenService {
  return &CitizenService{rp}
}

func (cs *CitizenService) ListAllCitizens() ([]*domain.Citizen, error) {
  return cs.rp.ListAllCitizens()
}

func (cs *CitizenService) GetCitizenById(id uint64) (*domain.Citizen, error) {
  return cs.rp.GetCitizenById(id)
}
