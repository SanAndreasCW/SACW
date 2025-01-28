package dev

import "github.com/kodeyeen/omp"

func init() {
	omp.Events.Listen(omp.EventTypePlayerClickMap, onPlayerClickMap)
	omp.Commands.Add("position", getPosition)
	omp.Commands.Add("spawn", spawnVehicle)
}
