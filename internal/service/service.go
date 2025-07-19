package service

import (
	"database/sql"
	"net/url"
	"url-shorter/config"
	"url-shorter/internal/checker"
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
	if !checker.ExistUrl(req.LongUrl) {
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
