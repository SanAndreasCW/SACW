package company

import (
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
)

func init() {
	// Auth Events
	omp.ListenFunc(auth.EventTypeOnAuthSuccess, onAuthSuccess)
	// Default Events
	omp.ListenFunc(omp.EventTypeGameModeInit, onGameModeInit)
	omp.ListenFunc(omp.EventTypeGameModeExit, onGameModeExit)
	omp.ListenFunc(omp.EventTypePlayerKeyStateChange, onPlayerKeyStateChange)
	// Commands
	//omp.Commands.Add("companies", companiesCommand)
	//omp.Commands.Add("company", companyCommand)
	// Module Announcement
	logger.Info("Company module initialized")
}
