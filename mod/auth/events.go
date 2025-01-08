package auth

import (
	"github.com/SanAndreasCW/SACW/mod/types"
	"github.com/kodeyeen/event"
)

var (
	Events = event.NewDispatcher()
)

const (
	EventTypeOnAuthSuccess event.Type = "onAuthSuccess"
)

type OnAuthSuccessEvent struct {
	PlayerI *types.PlayerI
	Success bool
}
