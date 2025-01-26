package job

import (
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"slices"
)

func onGameModeInit(e *omp.GameModeInitEvent) bool {
	commons.Jobs[enums.Delivery] = &commons.Job{
		ID:     enums.Delivery,
		Name:   "Delivery",
		Payout: 1000,
		VehicleModels: []omp.VehicleModel{
			omp.VehicleModelPicador,
		},
		CheckpointLocations: []omp.Vector3{
			omp.Vector3{X: 2233.792236, Y: -2216.103516, Z: 13.546875},
		},
	}
	logger.Info("Job module initialized.")
	return true
}

func onPlayerStateChange(e *omp.PlayerStateChangeEvent) bool {
	playerI := commons.PlayersI[e.Player.ID()]
	playerVehicle, _ := playerI.Vehicle()
	if playerI.Job == nil {
		return true
	}
	if playerI.Job.OnDuty == false {
		return true
	}
	if playerI.Job.Idle == true {
		if e.OldState == omp.PlayerStateOnFoot && (e.NewState == omp.PlayerStateDriver || e.NewState == omp.PlayerStatePassenger) {
			if !slices.Contains(playerI.Job.Job.VehicleModels, playerVehicle.Model()) == true {
				return true
			}
			playerI.Job.Idle = false
			if playerI.Job.Checkpoint != nil {
				playerI.Job.Checkpoint.Enable()
			}
		}
	} else {
		if e.OldState == omp.PlayerStateDriver && e.NewState == omp.PlayerStateOnFoot {
			playerI.Job.Idle = true
			if playerI.Job.Checkpoint != nil {
				playerI.Job.Checkpoint.Disable()
			}
		}
	}
	return true
}
