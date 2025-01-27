package job

import "github.com/RahRow/omp"

func init() {
	// Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerStateChange, onPlayerStateChange)
	omp.Events.Listen(omp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
}
