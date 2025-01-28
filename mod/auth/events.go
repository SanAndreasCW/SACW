package auth

import (
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/kodeyeen/event"
	"github.com/kodeyeen/omp"
)

var (
	Events = event.NewDispatcher()
)

const (
	EventTypeOnAuthSuccess omp.EventType = "onAuthSuccess"
)

type OnAuthSuccessEvent struct {
	PlayerI *commons.PlayerI
	Success bool
}
