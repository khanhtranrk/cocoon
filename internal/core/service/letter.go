package service

import (
	"github.com/khanhtranrk/cocoon/internal/core/domain"
	"github.com/khanhtranrk/cocoon/internal/core/port"
)

type LetterService struct {
  lr port.LetterRepository
}

func NewLetterService(lr port.LetterRepository) *LetterService {
  return &LetterService{lr}
}

func (ls *LetterService) CreateSuspiciousLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "suspicious_letters")
}

func (ls *LetterService) CreateProccessedLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "response_later_letters")
}

func (ls *LetterService) CreateSentLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "sent_letters")
}

func (ls *LetterService) DeleteSuspiciousLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "suspicious_letters")
}

func (ls *LetterService) DeleteProccessedLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "response_later_letters")
}

func (ls *LetterService) DeleteSentLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "sent_letters")
}

func (ls *LetterService) UpdateSuspiciousLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "suspicious_letters")
}

func (ls *LetterService) UpdateProccessedLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "response_later_letters")
}

func (ls *LetterService) UpdateSentLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "sent_letters")
}

func (ls *LetterService) GetSuspiciousLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "suspicious_letters")
}

func (ls *LetterService) GetProccessedLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "response_later_letters")
}

func (ls *LetterService) GetSentLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "sent_letters")
}

