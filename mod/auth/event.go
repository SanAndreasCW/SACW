package auth

import (
	"github.com/LosantosGW/go_LSGW/mod/types"
	"github.com/kodeyeen/event"
)

const (
	EventTypeOnAuthSuccess event.Type = "onAuthSuccess"
)

type OnAuthSuccessEvent struct {
	PlayerI *types.PlayerI
	Success bool
}
