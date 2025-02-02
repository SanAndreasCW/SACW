package job

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
)

func OnGameModeInit(_ context.Context, _ omp.Event) error {
	deliveryCheckpoints := []*omp.Vector3{
		{X: 2233.792236, Y: -2216.103516, Z: 13.546875},
	}
	deliveryLookups := []*omp.Vector3{
		{X: 2181.576660, Y: -2302.060547, Z: 13.546875},
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
		Group:               enums.CheckpointVehicle,
	}
	logger.Info("Job module initialized.")
	return nil
}

func onPlayerStateChange(_ context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerStateChangeEvent)
	playerI := commons.PlayersI[ep.Player.ID()]
	if playerI.Job == nil {
		return nil
	}
	if playerI.Job.OnDuty == false {
		return nil
	}
	if playerI.Job.Idle == true {
		if ep.OldState == omp.PlayerStateOnFoot && (ep.NewState == omp.PlayerStateDriver || ep.NewState == omp.PlayerStatePassenger) {
			playerI.SetJobCheckpoint()
		}
	} else {
		if ep.OldState == omp.PlayerStateDriver && ep.NewState == omp.PlayerStateOnFoot {
			playerI.Job.Idle = true
			if playerI.Job != nil {
				playerI.DefaultCheckpoint().Disable()
			}
		}
	}
	return nil
}

func onPlayerEnterCheckpoint(_ context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerEnterCheckpointEvent)
	playerI := commons.PlayersI[ep.Player.ID()]
	if playerI.Job == nil {
		return nil
	}
	if !playerI.Job.OnDuty {
		return nil
	}
	if playerI.Job.Idle {
		return nil
	}
	playerI.DefaultCheckpoint().Disable()
	if playerI.Job.Cargo.Loaded {
		playerI.SendClientMessage("You've successfully delivered the cargo.", colors.SuccessColor.Hex)
		playerI.Job.Cargo.Amount -= 1
		playerI.GiveMoney(int32(playerI.Job.Cargo.Value * 1000))

		if playerI.Job.Cargo.Amount <= 0 {
			playerI.Job.Cargo.Loaded = false
			playerI.SendClientMessage("You may need to reload your cargo at lookup point.", colors.NoteColor.Hex)
		}
	} else {
		playerI.Job.Cargo.Amount = 1
		playerI.Job.Cargo.Loaded = true
		playerI.SendClientMessage("You've successfully loaded the cargo.", colors.SuccessColor.Hex)
		playerI.SendClientMessage("You've to deliver cargo to the targeted location.", colors.NoteColor.Hex)
	}
	playerI.SetJobCheckpoint()
	return nil
}

func onAuthSuccess(ctx context.Context, e omp.Event) error {
	ep := e.Payload().(*auth.OnAuthSuccessEvent)
	playerI := commons.PlayersI[ep.PlayerI.ID()]
	playerI.LoadJobsInfo(ctx)
	return nil
}
