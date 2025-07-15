package service

import (
	"database/sql"
	"net/url"
	"url-shorter/config"
	"url-shorter/internal/errs"
	"url-shorter/internal/models"
	"url-shorter/internal/repo"
	"url-shorter/internal/shorter"
)

type Service interface {
	AddUrl(req *models.Requset) (models.Response, error)
	FindUrl(req *models.Requset) (models.Response, error)
}

type service struct {
	repo repo.Repo
}

func NewService(db *sql.DB, cfg *config.Config) Service {
	return service{repo: repo.NewRepo(db, cfg)}
}

func (this service) AddUrl(req *models.Requset) (models.Response, error) {
	if _, err := url.ParseRequestURI(req.LongUrl); err != nil {
		return models.Response{}, errs.InvalidURL()
	}

	resp := models.Response{
		ShortUrl: shorter.ShortenUrl(req.LongUrl),
		LongUrl:  req.LongUrl,
	}

	err := this.repo.AddUrl(&resp)
	if err != nil {
		return models.Response{}, errs.InvalidSQL("can't add new url")
	}

	return resp, nil
}

func (this service) FindUrl(req *models.Requset) (models.Response, error) {
	resp, err := this.repo.FindShortUrl(req.LongUrl)
	if err != nil {
		return models.Response{}, errs.InvalidSQL("can't find new url")
	}

	return resp, nil
}
