package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage struct {
	config          *Config
	db              *sql.DB
	songsRepositiry *SongsRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}
func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	storage.db = db
	log.Println("Database connection created successfuly!")
	return nil
}
func (storage *Storage) Close() {
	storage.db.Close()
}
func (s *Storage) Songs() *SongsRepository {
	if s.songsRepositiry != nil {
		return s.songsRepositiry
	}
	s.songsRepositiry = &SongsRepository{
		storage: s,
	}
	return s.songsRepositiry
}
