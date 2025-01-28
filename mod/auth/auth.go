package auth

import (
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
)

func init() {
	// Default Events
	omp.ListenFunc(omp.EventTypeGameModeInit, onGameModeInit)
	omp.ListenFunc(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.ListenFunc(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.ListenFunc(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.ListenFunc(omp.EventTypePlayerDisconnect, onPlayerDisconnect)
	omp.ListenFunc(omp.EventTypePlayerText, onPlayerText)
	// Application Initiation Complete
	logger.Info("Auth Module Initialized")
}
