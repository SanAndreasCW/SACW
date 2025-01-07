package auth

import (
	"github.com/LosantosGW/go_LSGW/mod/logger"
	"github.com/RahRow/omp"
)

func init() {
	// Default Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.Events.Listen(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.Events.Listen(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.Events.Listen(omp.EventTypePlayerDisconnect, onPlayerDisconnect)

	// Application Initiation Complete
	logger.Info("Auth Module Initialized")
}
