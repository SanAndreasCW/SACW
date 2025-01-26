package main

import (
	"github.com/RahRow/omp"
	_ "github.com/SanAndreasCW/SACW/mod/company"
	"github.com/SanAndreasCW/SACW/mod/database"
	_ "github.com/SanAndreasCW/SACW/mod/dev"
	_ "github.com/SanAndreasCW/SACW/mod/job"
	"github.com/SanAndreasCW/SACW/mod/logger"
	_ "github.com/joho/godotenv"
)

func init() {
	omp.Events.Listen(omp.EventTypeGameModeExit, func(event *omp.GameModeExitEvent) bool {
		logger.Info("Database connection closed")
		err := database.DB.Close()
		if err != nil {
			logger.Error("Failed to close database connection: %s", err)
		}
		return true
	})
}

func main() {}
