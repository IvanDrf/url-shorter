package repo

import (
	"database/sql"
	"fmt"
	"url-shorter/config"
	"url-shorter/internal/models"
)

type Repo interface {
	AddUrl(req *models.Response) error

	FindShortUrl(long string) (models.Response, error)
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

func (this repo) FindShortUrl(long string) (models.Response, error) {
	res, _ := this.db.Query(fmt.Sprintf("SELECT short_url FROM %s.%s WHERE long_url = '%s'", this.dbName, this.dbName, long))
	resp := models.Response{}

	err := res.Scan(&resp.ShortUrl, &resp.LongUrl)
	return resp, err
}
