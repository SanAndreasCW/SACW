package timer

import "time"

func setTimeout(timeout *Timeout) *time.Timer {
	duration := time.Duration(timeout.LifeTime) * time.Millisecond
	return time.AfterFunc(duration, timeout.Callback)
}
