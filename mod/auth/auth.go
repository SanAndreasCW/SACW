package auth

import (
	"github.com/SanAndreasCW/SACW/mod/logger"
	
)

func init() {
	// Default Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.Events.Listen(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.Events.Listen(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.Events.Listen(omp.EventTypePlayerDisconnect, onPlayerDisconnect)
	omp.Events.Listen(omp.EventTypePlayerText, onPlayerText)
	// Application Initiation Complete
	logger.Info("Auth Module Initialized")
}
