package cmd

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/kodeyeen/omp"
)

func onPlayerCommandText(ctx context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerCommandTextEvent)
	if Commands.Has(ep.Name) {
		Commands.run(ep.Name, &Command{
			Name:     ep.Name,
			Args:     ep.Args,
			Sender:   ep.Sender,
			RawValue: ep.RawValue,
		})
	}
	return commons.NewRefuse()
}
