package job

import "github.com/kodeyeen/omp"

func init() {
	// Events
	omp.ListenFunc(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerStateChange, onPlayerStateChange)
	omp.Events.Listen(omp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
}
