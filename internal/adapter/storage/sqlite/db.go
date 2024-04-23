package sqlite

import (
	"database/sql"

	"github.com/khanhtranrk/cocoon/internal/adapter/config"
	_ "github.com/mattn/go-sqlite3"
)

func New(config *config.Config) (*sql.DB, error) {
  db, err := sql.Open("sqlite3", config.DatabaseUrl)

  if err != nil {
    return nil, err
  }

  return db, nil
}

