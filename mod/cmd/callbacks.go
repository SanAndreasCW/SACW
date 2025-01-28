package cmd

import (
	"context"
	"fmt"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/kodeyeen/omp"
)

func onPlayerCommandText(_ context.Context, e omp.Event) error {
	ep, ok := e.Payload().(*omp.PlayerCommandTextEvent)
	if !ok {
		return commons.NewRefuse()
	}
	if Commands.Has(ep.Name) {
		Commands.run(ep.Name, &Command{
			Name:     ep.Name,
			Args:     ep.Args,
			Sender:   ep.Sender,
			RawValue: ep.RawValue,
		})
	} else {
		ep.Sender.SendClientMessage("Error: Command not found", colors.ErrorHex)
	}
	return nil
}

func onPlayerText(ctx context.Context, e omp.Event) error {
	ep, ok := e.Payload().(*omp.PlayerTextEvent)
	if ok {
		msg := fmt.Sprintf("[ID:%d|Name:%s]: %s", ep.Player.ID(), ep.Player.Name(), ep.Message)
		commons.SendClientMessageToAll(msg, colors.WhiteHex)
	}
	return commons.NewRefuse()
}
