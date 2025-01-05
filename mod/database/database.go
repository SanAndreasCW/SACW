package database

import (
	"context"

	"github.com/LosantosGW/go_LSGW/mod/logger"
	"github.com/RahRow/omp"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, func(e *omp.GameModeInitEvent) bool {
		DB, err := pgx.Connect(context.Background(), "user=postgres password=dev host=localhost port=5432 dbname=lsgw")

		if err != nil {
			logger.Error("Failed to connect to database: %s", err)
			omp.SendRCONCommand("exit")
			return true
		}

		if err := DB.Ping(context.Background()); err != nil {
			logger.Error("Failed to ping database: %s", err)
			omp.SendRCONCommand("exit")
			return true
		}

		logger.Info("Database module initialized")
		return true
	})

	omp.Events.Listen(omp.EventTypeGameModeExit, func(event *omp.GameModeExitEvent) bool {
		DB.Close(context.Background())
		return true
	})
}
