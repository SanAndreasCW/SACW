package cmd

import (
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
)

var Commands = newCommandManager()

func init() {
	omp.ListenFunc(omp.EventTypePlayerCommandText, onPlayerCommandText)
	omp.ListenFunc(omp.EventTypePlayerText, onPlayerText)

	logger.Info("Command utility loaded")
}
