package store

import (
	"database/sql"
	"github.com/wspectra/urlShort/internal/config"

	_ "github.com/lib/pq"
)

type Postgres struct {
	conf *config.Config
	db   *sql.DB
}

func NewPstStore(conf *config.Config) *Postgres {
	return &Postgres{conf: conf}
}

func (p *Postgres) Open() error {
	db, err := sql.Open("postgres", p.conf.DatabaseUrl)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	p.db = db
	return nil
}

func (p *Postgres) Close() {
	p.db.Close()
}

func (p *Postgres) GetInfo(find string) (string, error) {
	return "", nil
}

func (p *Postgres) PostInfo(info string) (string, error) {
	return "", nil
}
