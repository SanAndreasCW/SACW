package job

import "github.com/kodeyeen/omp"

func init() {
	// Events
	omp.ListenFunc(omp.EventTypePlayerStateChange, onPlayerStateChange)
	omp.ListenFunc(omp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
}
