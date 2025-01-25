package job

import (
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
)

func onGameModeInit(e *omp.GameModeInitEvent) bool {
	commons.Jobs[enums.Delivery] = &commons.Job{
		ID:     enums.Delivery,
		Name:   "Delivery",
		Payout: 1000,
	}
	logger.Info("Job module initialized.")
	return true
}
