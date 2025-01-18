package database

import (
	"database/sql"
	_ "database/sql/driver"

	_ "github.com/lib/pq"

	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/logger"
)

var DB *sql.DB

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, func(e *omp.GameModeInitEvent) bool {
		var err error = nil
		DB, err = sql.Open("postgres", "user=postgres password=dev host=localhost port=5432 dbname=sacw sslmode=disable")
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

}
