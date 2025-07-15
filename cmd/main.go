package main

import (
	"url-shorter/config"
	"url-shorter/internal/database"
	"url-shorter/internal/transport/server"
	"url-shorter/logger"
)

func main() {
	cfg := config.InitCFG()

	logger := logger.InitLogger(cfg)
	server := server.NewServer(logger)
	server.RegisterRoutes()

	db := database.InitDB(cfg)
	defer db.Close()

	server.Start(cfg)
}
