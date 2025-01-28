package dev

import (
	"github.com/SanAndreasCW/SACW/mod/cmd"
	"github.com/kodeyeen/omp"
)

func init() {
	omp.ListenFunc(omp.EventTypePlayerClickMap, onPlayerClickMap)
	cmd.Commands.Add("position", getPosition)
	cmd.Commands.Add("spawn", spawnVehicle)
}
