package handlers

import (
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"url-shorter/config"
	"url-shorter/internal/errs"
	"url-shorter/internal/models"
	"url-shorter/internal/service"
)

type Handler interface {
	PostHandler(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.Service
	logger  *slog.Logger
}

func NewHandler(cfg *config.Config, db *sql.DB, logger *slog.Logger) Handler {
	return handler{logger: logger, service: service.NewService(db, cfg)}
}

func (this handler) PostHandler(w http.ResponseWriter, r *http.Request) {
	this.logger.Info("POST req")
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(errs.InvalidMediaType().Error()))

		this.logger.Error("invalid content type")
		return
	}

	req := models.Requset{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(errs.InvalidJSON().Error()))

		this.logger.Error("invalid json")
		return
	}

	resp, err := this.service.AddUrl(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))

		this.logger.Warn("url already in db")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

	this.logger.Info("POST success")
}

func (this handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	this.logger.Info("GET req")

	short := r.PathValue("short")
	resp, err := this.service.FindUrl(short)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(errs.InvalidShortURL().Error()))

		this.logger.Warn("url is not in db")
		return
	}

	this.logger.Info("GET success")
	http.Redirect(w, r, resp.LongUrl, http.StatusFound)
}
