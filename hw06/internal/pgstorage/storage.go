package pgstorage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	config *Config
	db     *sqlx.DB
}

func New(config *Config) *Storage {
	return &Storage{config: config}
}

func (s *Storage) Open() error {
	db, err := sqlx.Connect("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db
	return nil
}

func (s *Storage) Close() {
	_ = s.db.Close()
}

func (s *Storage) Users() *UserRepo {
	return &UserRepo{storage: s}
}
