package main

import (
	"github.com/kodeyeen/omp"
)

func init() {
	omp.Events.Listen(omp.EventTypePlayerConnect, func(e *omp.PlayerConnectEvent) bool {
		e.Player.SendClientMessage("Hello, world!", 0xFFFF00FF)
		return true
	})
}

func main() {

}
