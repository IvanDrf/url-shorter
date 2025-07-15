package database

import (
	"database/sql"
	"fmt"
	"log"
	"url-shorter/config"
	"url-shorter/internal/errs"

	_ "github.com/go-sql-driver/mysql"
)

/*dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)*/

func InitDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(errs.InvalidDBConnection())
	}

	return db
}
