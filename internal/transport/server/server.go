package server

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"url-shorter/config"
	"url-shorter/internal/errs"
	"url-shorter/internal/transport/handlers"
)

type Server struct {
	server   *http.ServeMux
	handlers handlers.Handler
}

func NewServer(logger *slog.Logger) Server {
	return Server{
		server:   http.NewServeMux(),
		handlers: handlers.NewHandler(logger),
	}
}

func (this *Server) Start(cfg *config.Config) {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.ServerPort), this.server); err != nil {
		log.Fatal(errs.InvalidStart())
	}
}

func (this *Server) RegisterRoutes() {
	this.server.HandleFunc("POST /urls", this.handlers.PostHandler)
	this.server.HandleFunc("GET /urls", this.handlers.GetHandler)
}
