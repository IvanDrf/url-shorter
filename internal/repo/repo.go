package repo

import (
	"database/sql"
	"fmt"
	"url-shorter/config"
	"url-shorter/internal/models"
)

type Repo interface {
	AddUrl(req *models.Response) error

	FindLongUrl(short string) (models.Response, error)
}

type repo struct {
	db     *sql.DB
	dbName string
}

func NewRepo(db *sql.DB, cfg *config.Config) Repo {
	return repo{db: db, dbName: cfg.DBName}
}

func (this repo) AddUrl(req *models.Response) error {
	_, err := this.db.Exec(fmt.Sprintf("INSERT INTO %s.%s (short_url, long_url) VALUES ('%s', '%s');", this.dbName, this.dbName, req.ShortUrl, req.LongUrl))
	return err
}

func (this repo) FindLongUrl(short string) (models.Response, error) {
	res := this.db.QueryRow(fmt.Sprintf("SELECT long_url FROM %s.%s WHERE short_url = '%s';", this.dbName, this.dbName, short))
	resp := models.Response{}

	err := res.Scan(&resp.LongUrl)

	return resp, err
}
