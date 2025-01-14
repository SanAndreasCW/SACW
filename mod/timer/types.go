package timer

type Timer struct {
	Callback func()
	TickRate int
	Async    bool
}

type Timeout struct {
	Callback func()
	LifeTime int
}
