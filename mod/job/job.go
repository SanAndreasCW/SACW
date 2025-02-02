package job

import (
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/kodeyeen/omp"
)

func init() {
	// Events
	omp.ListenFunc(omp.EventTypePlayerStateChange, onPlayerStateChange)
	omp.ListenFunc(omp.EventTypePlayerEnterCheckpoint, onPlayerEnterCheckpoint)
	omp.ListenFunc(auth.EventTypeOnAuthSuccess, onAuthSuccess)
}
