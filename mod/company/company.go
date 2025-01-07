package company

import (
	"github.com/LosantosGW/go_LSGW/mod/auth"
	"github.com/RahRow/omp"
)

func init() {
	// Auth Events
	auth.Events.Listen(auth.EventTypeOnAuthSuccess, onAuthSuccess)

	// Default Events
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)

	// Commands
	omp.Commands.Add("companies", companyCommand)
}
