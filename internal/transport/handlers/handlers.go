package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errs.InvalidMediaType())

		this.logger.Error("invalid content type")
		return
	}

	req := models.Requset{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errs.InvalidJSON())

		this.logger.Error("invalid json")
		return
	}

	resp, err := this.service.AddUrl(&req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errs.InvalidSQL("cant add new url"))

		if errors.Is(err, errs.InvalidURL()) {
			this.logger.Warn("invalid url")
		} else {
			this.logger.Warn("cant add new url")
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

	this.logger.Info("POST success")
}

func (this handler) GetHandler(w http.ResponseWriter, r *http.Request) {
	this.logger.Info("GET req")

	short := r.PathValue("short")
	resp, err := this.service.FindUrl(short)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errs.InvalidShortURL())

		this.logger.Warn("url is not in db")
		return
	}

	this.logger.Info("GET success")
	http.Redirect(w, r, resp.LongUrl, http.StatusFound)
}
