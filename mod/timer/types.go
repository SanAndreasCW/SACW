package timer

import "time"

type Timer struct {
	Callback func()
	Duration time.Duration
	Async    bool
}

type Timeout struct {
	Callback func()
	LifeTime int
}
