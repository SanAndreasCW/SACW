package job

import "github.com/RahRow/omp"

func init() {
	// Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
}
