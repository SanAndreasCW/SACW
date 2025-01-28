package main

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/database"
	_ "github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"

	_ "github.com/SanAndreasCW/SACW/mod/company"
	_ "github.com/SanAndreasCW/SACW/mod/dev"
	_ "github.com/SanAndreasCW/SACW/mod/job"
	_ "github.com/joho/godotenv"
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
