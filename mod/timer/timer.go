package timer

import "time"

func SetTimer(timer *Timer) chan bool {
	ticker := time.NewTicker(timer.Duration)
	clearTimer := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				if timer.Async {
					go timer.Callback()
				} else {
					timer.Callback()
				}
			case <-clearTimer:
				ticker.Stop()
				return
			}
		}
	}()
	return clearTimer
}
