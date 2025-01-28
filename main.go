package main

import (
	"context"
	"database/sql"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/company"
	"github.com/SanAndreasCW/SACW/mod/database"
	_ "github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/job"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"

	_ "github.com/SanAndreasCW/SACW/mod/company"
	_ "github.com/SanAndreasCW/SACW/mod/dev"
	_ "github.com/SanAndreasCW/SACW/mod/job"
	_ "github.com/joho/godotenv"
)

func init() {
	omp.ListenFunc(omp.EventTypeGameModeInit, func(ctx context.Context, e omp.Event) error {
		var err error = nil
		database.DB, err = sql.Open("postgres", "user=postgres password=dev host=localhost port=5432 dbname=sacw sslmode=disable")
		if err != nil {
			logger.Error("Failed to connect to database: %s", err)
			omp.SendRCONCommand("exit")
			return nil
		}
		if err := database.DB.Ping(); err != nil {
			logger.Error("Failed to ping database: %s", err)
			omp.SendRCONCommand("exit")
			return nil
		}
		logger.Info("Database module initialized")
		err = auth.OnGameModeInit(ctx, e)
		err = company.OnGameModeInit(ctx, e)
		err = job.OnGameModeInit(ctx, e)
		return err
	})
	omp.ListenFunc(omp.EventTypeGameModeExit, func(ctx context.Context, event omp.Event) error {
		err := company.OnGameModeExit(ctx, event)
		logger.Info("Database connection closed")
		defer func(DB *sql.DB) {
			err := DB.Close()
			if err != nil {
				logger.Error("Failed to close database connection: %s", err)
			}
		}(database.DB)
		return err
	})
}

func main() {}
