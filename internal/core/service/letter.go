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

func (ls *LetterService) CreateHandleLaterLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "handle_later_letters")
}

func (ls *LetterService) CreateResponseLaterLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "response_later_letters")
}

func (ls *LetterService) CreateSentLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "sent_letters")
}

func (ls *LetterService) CreateReceivedLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.CreateLetter(letter, "received_letters")
}

func (ls *LetterService) DeleteHandleLaterLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "handle_later_letters")
}

func (ls *LetterService) DeleteResponseLaterLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "response_later_letters")
}

func (ls *LetterService) DeleteSentLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "sent_letters")
}

func (ls *LetterService) DeleteReceivedLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.DeleteLetter(letter, "received_letters")
}

func (ls *LetterService) UpdateHandleLaterLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "handle_later_letters")
}

func (ls *LetterService) UpdateResponseLaterLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "response_later_letters")
}

func (ls *LetterService) UpdateSentLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "sent_letters")
}

func (ls *LetterService) UpdateReceivedLetter(letter *domain.Letter) (*domain.Letter, error) {
  return ls.lr.UpdateLetter(letter, "received_letters")
}

func (ls *LetterService) GetHandleLaterLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "handle_later_letters")
}

func (ls *LetterService) GetResponseLaterLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "response_later_letters")
}

func (ls *LetterService) GetSentLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "sent_letters")
}

func (ls *LetterService) GetReceivedLetterById(id uint64) (*domain.Letter, error) {
  return ls.lr.GetLetterById(id, "received_letters")
}
