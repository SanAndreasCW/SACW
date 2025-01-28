package main

import (
	"context"
	_ "github.com/SanAndreasCW/SACW/mod/company"
	"github.com/SanAndreasCW/SACW/mod/database"
	_ "github.com/SanAndreasCW/SACW/mod/dev"
	_ "github.com/SanAndreasCW/SACW/mod/job"
	"github.com/SanAndreasCW/SACW/mod/logger"
	_ "github.com/joho/godotenv"
	"github.com/kodeyeen/omp"
)

func init() {
	omp.ListenFunc(omp.EventTypeGameModeExit, func(ctx context.Context, event omp.Event) error {
		logger.Info("Database connection closed")
		err := database.DB.Close()
		if err != nil {
			logger.Error("Failed to close database connection: %s", err)
		}
		return nil
	})
}

func main() {}
