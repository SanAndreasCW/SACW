package job

import (
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	
)

func onGameModeInit(e *omp.GameModeInitEvent) bool {
	deliveryCheckpoints := []*omp.Vector3{
		&omp.Vector3{X: 2233.792236, Y: -2216.103516, Z: 13.546875},
	}
	deliveryLookups := []*omp.Vector3{
		&omp.Vector3{X: 2181.576660, Y: -2302.060547, Z: 13.546875},
	}
	commons.Jobs[enums.Delivery] = &commons.Job{
		ID:     enums.Delivery,
		Name:   "Delivery",
		Payout: 1000,
		VehicleModels: []omp.VehicleModel{
			omp.VehicleModelPicador,
			omp.VehicleModelMule,
			omp.VehicleModelPony,
			omp.VehicleModelBurrito,
			omp.VehicleModelBobcat,
			omp.VehicleModelRumpo,
			omp.VehicleModelYankee,
			omp.VehicleModelWalton,
			omp.VehicleModelBenson,
			omp.VehicleModelSadler,
			omp.VehicleModelBoxville,
			omp.VehicleModelBoxburg,
		},
		CheckpointLocations: deliveryCheckpoints,
		LookupLocations:     deliveryLookups,
	}
	logger.Info("Job module initialized.")
	return true
}

func onPlayerStateChange(e *omp.PlayerStateChangeEvent) bool {
	playerI := commons.PlayersI[e.Player.ID()]
	if playerI.Job == nil {
		return true
	}
	if playerI.Job.OnDuty == false {
		return true
	}
	if playerI.Job.Idle == true {
		if e.OldState == omp.PlayerStateOnFoot && (e.NewState == omp.PlayerStateDriver || e.NewState == omp.PlayerStatePassenger) {
			playerI.SetJobCheckpoint()
		}
	} else {
		if e.OldState == omp.PlayerStateDriver && e.NewState == omp.PlayerStateOnFoot {
			playerI.Job.Idle = true
			if playerI.Job != nil {
				playerI.DefaultCheckpoint().Disable()
			}
		}
	}
	return true
}

func onPlayerEnterCheckpoint(e *omp.PlayerEnterCheckpointEvent) bool {
	playerI := commons.PlayersI[e.Player.ID()]
	if playerI.Job == nil {
		return true
	}
	if !playerI.Job.OnDuty {
		return true
	}
	if playerI.Job.Idle {
		return true
	}
	playerI.DefaultCheckpoint().Disable()
	if playerI.Job.Cargo.Loaded {
		playerI.SendClientMessage("You've successfully delivered the cargo.", colors.SuccessHex)
		playerI.Job.Cargo.Amount -= 1
		playerI.GiveMoney(int32(playerI.Job.Cargo.Value * 1000))

		if playerI.Job.Cargo.Amount <= 0 {
			playerI.Job.Cargo.Loaded = false
			playerI.SendClientMessage("You may need to reload your cargo at lookup point.", colors.NoteHex)
		}
	} else {
		playerI.Job.Cargo.Amount = 1
		playerI.Job.Cargo.Loaded = true
		playerI.SendClientMessage("You've successfully loaded the cargo.", colors.SuccessHex)
		playerI.SendClientMessage("You've to deliver cargo to the targeted location.", colors.NoteHex)
	}
	playerI.SetJobCheckpoint()
	return true
}
