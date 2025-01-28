package dev

import (
	"context"
	"github.com/kodeyeen/omp"
)

func onPlayerClickMap(ctx context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerClickMapEvent)
	ep.Player.SetPosition(omp.Vector3{
		X: ep.Position.X,
		Y: ep.Position.Y,
		Z: ep.Position.Z,
	})
	return nil
}
