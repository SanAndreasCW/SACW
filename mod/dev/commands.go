package dev

import (
	"fmt"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/logger"
)

func getPosition(cmd *omp.Command) {
	player := cmd.Sender
	position := player.Position()
	player.SendClientMessage(
		fmt.Sprintf("position X: %f | Y: %f | Z: %f", position.X, position.Y, position.Z),
		colors.NoticeHex,
	)
	logger.Debug("position X: %f | Y: %f | Z: %f", position.X, position.Y, position.Z)
}
