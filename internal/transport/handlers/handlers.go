package handlers

import (
	"log/slog"
	"net/http"
)

type Handler interface {
	PostHandler(w http.ResponseWriter, r *http.Request)
	GetHandler(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger *slog.Logger
}

func NewHandler(logger *slog.Logger) Handler {
	return handler{logger: logger}
}

func (this handler) PostHandler(w http.ResponseWriter, r *http.Request) {

}

func (this handler) GetHandler(w http.ResponseWriter, r *http.Request) {

}
