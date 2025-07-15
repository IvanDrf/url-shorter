package main

import (
	"fmt"
	"url-shorter/config"
	"url-shorter/internal/database"
	"url-shorter/logger"
)

func main() {
	cfg := config.InitCFG()

	fmt.Println(*cfg)
	logger := logger.InitLogger(cfg)
	db := database.InitDB(cfg)
	defer db.Close()

	logger.Info("test")

}
