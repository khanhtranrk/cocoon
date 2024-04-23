package port

import "github.com/khanhtranrk/cocoon/internal/core/domain"

type LetterRepository interface {
  CreateLetter(letter *domain.Letter, tableName string) (*domain.Letter, error)
  UpdateLetter(letter *domain.Letter, tableName string) (*domain.Letter, error)
  DeleteLetter(letter *domain.Letter, tableName string) (*domain.Letter, error)
  GetLetterById(id uint64, tableName string) (*domain.Letter, error)
}

type LetterService interface {
  CreateUnprocessedLetter(letter *domain.Letter) (*domain.Letter, error)
  CreateUnsentLaterLetter(letter *domain.Letter) (*domain.Letter, error)
  CreateProcessedLetter(letter *domain.Letter) (*domain.Letter, error)
  CreateSentLetter(letter *domain.Letter) (*domain.Letter, error)

  DeleteUnprocessedLetter(letter *domain.Letter) (*domain.Letter, error)
  DeleteUnsentLetter(letter *domain.Letter) (*domain.Letter, error)
  DeleteProcessedLetter(letter *domain.Letter) (*domain.Letter, error)
  DeleteSentLetter(letter *domain.Letter) (*domain.Letter, error)

  UpdateUnprocessedLetter(letter *domain.Letter) (*domain.Letter, error)
  UpdateUnsentLetter(letter *domain.Letter) (*domain.Letter, error)
  UpdateProcessedLetter(letter *domain.Letter) (*domain.Letter, error)
  UpdateSentLetter(letter *domain.Letter) (*domain.Letter, error)

  GetUnprocessedLetterById(id string) (*domain.Letter, error)
  GetUnsentLetterById(id string) (*domain.Letter, error)
  GetProcessedLetterById(id string) (*domain.Letter, error)
  GetSentLetterById(id string) (*domain.Letter, error)
}

