package store

import (
	"database/sql"
	"errors"
	"github.com/wspectra/urlShort/internal/config"
	"github.com/wspectra/urlShort/internal/pkg/utils"

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

func (p *Postgres) GetInfo(find string) (string, error) {
	var (
		long_url  string
		short_url string
	)
	if err := p.db.QueryRow(
		"SELECT long_url, short_url FROM urls where short_url = $1",
		find).Scan(&long_url, &short_url); err != nil {
		return "", errors.New("longUrl not found")
	}
	return long_url, nil
}

func (p *Postgres) PostInfo(info string) (string, error) {
	//проверка на наличие ссылки в базе
	var shortUrl string
	if err := p.db.QueryRow(
		"SELECT  short_url FROM urls where long_url = $1",
		info).Scan(&shortUrl); err == nil {
		return shortUrl, nil
	}

	sqlStatement := `
	INSERT INTO urls (long_url, short_url)
	VALUES ($1, $2)`
	randomString := utils.GenerateRandomString()
	_, err := p.db.Exec(sqlStatement, info, randomString)
	if err != nil {
		return "", err
	}
	return randomString, nil
}
