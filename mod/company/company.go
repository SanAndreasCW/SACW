package company

import (
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
)

func init() {
	// Auth Events
	auth.Events.Listen(auth.EventTypeOnAuthSuccess, onAuthSuccess)

	// Default Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)

	// Commands
	omp.Commands.Add("companies", companyCommand)
}
