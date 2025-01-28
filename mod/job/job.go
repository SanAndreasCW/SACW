package job

import "github.com/kodeyeen/omp"

func init() {
	// Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerStateChange, onPlayerStateChange)
	omp.Events.Listen(omp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
}
