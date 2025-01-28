package auth

import (
	"github.com/kodeyeen/omp"
)

func init() {
	// Default Events
	omp.ListenFunc(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.ListenFunc(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.ListenFunc(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.ListenFunc(omp.EventTypePlayerDisconnect, onPlayerDisconnect)
	omp.ListenFunc(omp.EventTypePlayerText, onPlayerText)
}
