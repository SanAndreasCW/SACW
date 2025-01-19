package dev

import "github.com/RahRow/omp"

func init() {
	omp.Events.Listen(omp.EventTypePlayerClickMap, onPlayerClickMap)

	omp.Commands.Add("position", getPosition)
}
