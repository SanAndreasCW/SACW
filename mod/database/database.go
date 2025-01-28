package database

import (
	"context"
	"database/sql"
	_ "database/sql/driver"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	omp.ListenFunc(omp.EventTypeGameModeInit, func(ctx context.Context, e omp.Event) error {
		var err error = nil
		DB, err = sql.Open("postgres", "user=postgres password=dev host=localhost port=5432 dbname=sacw sslmode=disable")
		if err != nil {
			logger.Error("Failed to connect to database: %s", err)
			omp.SendRCONCommand("exit")
			return nil
		}
		if err := DB.Ping(); err != nil {
			logger.Error("Failed to ping database: %s", err)
			omp.SendRCONCommand("exit")
			return nil
		}
		logger.Info("Database module initialized")
		return nil
	})
}
