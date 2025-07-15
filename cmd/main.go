package main

import (
	"url-shorter/config"
	"url-shorter/internal/database"
	"url-shorter/internal/transport/server"
	"url-shorter/logger"
)

func main() {
	cfg := config.InitCFG()

	db := database.InitDB(cfg)
	defer db.Close()

	logger := logger.InitLogger(cfg)

	server := server.NewServer(cfg, db, logger)
	server.RegisterRoutes()
	server.Start(cfg)
}
