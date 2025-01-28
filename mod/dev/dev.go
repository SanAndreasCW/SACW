package dev

import "github.com/kodeyeen/omp"

func init() {
	omp.ListenFunc(omp.EventTypePlayerClickMap, onPlayerClickMap)
	//omp.Commands.Add("position", getPosition)
	//omp.Commands.Add("spawn", spawnVehicle)
}
