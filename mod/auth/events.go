package auth

import (
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/kodeyeen/event"
)

var (
	Events = event.NewDispatcher()
)

const (
	EventTypeOnAuthSuccess event.Type = "onAuthSuccess"
)

type OnAuthSuccessEvent struct {
	PlayerI *commons.PlayerI
	Success bool
}
