package dev

import "github.com/RahRow/omp"

func onPlayerClickMap(e *omp.PlayerClickMapEvent) bool {
	e.Player.SetPosition(omp.Vector3{
		X: e.Position.X,
		Y: e.Position.Y,
		Z: e.Position.Z,
	})
	return true
}
