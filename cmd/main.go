package main

import (
	"url-shorter/config"
	"url-shorter/internal/logger"
)

func main() {
	cfg := config.InitCFG()
	logger := logger.InitLogger(cfg)

	logger.Info("test")

}
