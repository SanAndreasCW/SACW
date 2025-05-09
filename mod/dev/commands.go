package dev

import (
	"fmt"
	"github.com/SanAndreasCW/SACW/mod/cmd"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
)

func getPosition(cmd *cmd.Command) {
	player := cmd.Sender
	position := player.Position()
	player.SendClientMessage(
		fmt.Sprintf("position X: %f | Y: %f | Z: %f", position.X, position.Y, position.Z),
		colors.NoticeColor.Hex,
	)
	logger.Debug("position X: %f | Y: %f | Z: %f", position.X, position.Y, position.Z)
}

func spawnVehicle(cmd *cmd.Command) {
	player := cmd.Sender
	vehicleModel := cmd.Args[0]
	playerPos := player.Position()
	playerAngle := player.FacingAngle()
	vehicleID, err := commons.StringToInt[int32](&vehicleModel)
	if err != nil {
		player.SendClientMessage("[Spawn]: you've entered wrong vehicle id", colors.ErrorColor.Hex)
		return
	}
	veh, _ := omp.NewVehicle(omp.VehicleModel(vehicleID), omp.Vector3{
		X: playerPos.X,
		Y: playerPos.Y,
		Z: playerPos.Z,
	}, playerAngle)
	veh.PutPlayer(player, 0)
	player.SendClientMessage("[Spawn]: vehicle created", colors.SuccessColor.Hex)
}
