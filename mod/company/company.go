package company

import (
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
)

func init() {
	// Auth Events
	auth.Events.Listen(auth.EventTypeOnAuthSuccess, onAuthSuccess)
	// Default Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypeGameModeExit, onGameModeExit)
	omp.Events.Listen(omp.EventTypePlayerKeyStateChange, onPlayerKeyStateChange)
	// Commands
	omp.Commands.Add("companies", companiesCommand)
	omp.Commands.Add("company", companyCommand)
	// Module Announcement
	logger.Info("Company module initialized")
}
