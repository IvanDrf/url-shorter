package service

import (
	"context"
	"database/sql"
	"net/http"
	"net/url"
	"time"
	"url-shorter/config"
	"url-shorter/internal/errs"
	"url-shorter/internal/models"
	"url-shorter/internal/repo"
	"url-shorter/internal/shorter"
)

type Service interface {
	AddUrl(req *models.Requset) (models.Response, error)
	FindUrl(short string) (models.Response, error)
}

type service struct {
	repo    repo.Repo
	shorten shorter.Shorter
}

func NewService(db *sql.DB, cfg *config.Config) Service {
	return service{
		repo:    repo.NewRepo(db, cfg),
		shorten: shorter.NewShorten(),
	}
}

func (this service) AddUrl(req *models.Requset) (models.Response, error) {
	if !ExistUrl(req.LongUrl) {
		return models.Response{}, errs.InvalidURL()
	}

	if _, err := url.ParseRequestURI(req.LongUrl); err != nil {
		return models.Response{}, errs.InvalidURL()
	}

	resp := models.Response{
		ShortUrl: this.shorten.ShortenUrl(req.LongUrl),
		LongUrl:  req.LongUrl,
	}

	err := this.repo.AddUrl(&resp)
	return resp, err
}

func (this service) FindUrl(short string) (models.Response, error) {
	resp, err := this.repo.FindLongUrl(short)
	if err != nil {
		return models.Response{}, errs.InvalidSQL("can't find new url")
	}

	return resp, nil
}

func ExistUrl(url string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return false
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}

	defer resp.Body.Close()

	return resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusBadRequest

}
