package database

import (
	"database/sql"
	_ "database/sql/driver"

	_ "github.com/lib/pq"

	"github.com/LosantosGW/go_LSGW/mod/logger"
	"github.com/RahRow/omp"
)

var DB *sql.DB

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, func(e *omp.GameModeInitEvent) bool {
		var err error = nil
		DB, err = sql.Open("postgres", "user=postgres password=dev host=localhost port=5432 dbname=lsgw sslmode=disable")

		if err != nil {
			logger.Error("Failed to connect to database: %s", err)
			omp.SendRCONCommand("exit")
			return true
		}

		if err := DB.Ping(); err != nil {
			logger.Error("Failed to ping database: %s", err)
			omp.SendRCONCommand("exit")
			return true
		}

		logger.Info("Database module initialized")
		return true
	})

	omp.Events.Listen(omp.EventTypeGameModeExit, func(event *omp.GameModeExitEvent) bool {
		DB.Close()
		return true
	})
}
