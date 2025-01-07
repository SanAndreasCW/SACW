package auth

import (
	"context"
	"github.com/LosantosGW/go_LSGW/mod/setting"
	"github.com/kodeyeen/event"
	"time"

	"github.com/LosantosGW/go_LSGW/mod/database"
	"github.com/LosantosGW/go_LSGW/mod/logger"
	"github.com/LosantosGW/go_LSGW/mod/types"
	"github.com/RahRow/omp"
	"github.com/matthewhartstonge/argon2"
)

var (
	Events       = event.NewDispatcher()
	PlayersI     = make(map[int]*types.PlayerI, setting.MaxPlayers)
	PlayersCache = make(map[int]*types.PlayerCache, setting.MaxPlayers)
)

func onGameModeInit(_ *omp.GameModeInitEvent) bool {
	_, _ = omp.NewClass(255, 12, omp.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, 0.0, 0, 0, 0, 0, 0, 0)
	return true
}

func onPlayerConnect(e *omp.PlayerConnectEvent) bool {
	argon := argon2.DefaultConfig()
	player := e.Player
	playerCache := &types.PlayerCache{Player: player, LoginAttempts: 0}
	PlayersCache[player.ID()] = playerCache
	q := database.New(database.DB)

	ctx := context.Background()
	user, err := q.GetPlayerByUsername(ctx, player.Name())

	if err != nil {
		registerDialog := omp.NewInputDialog("Registration", "Please enter your password to register.", "Register", "Cancel")
		registerDialog.ShowFor(player)

		registerDialog.On(omp.EventTypeDialogResponse, func(e *omp.InputDialogResponseEvent) bool {
			if e.Response == omp.DialogResponseRight {
				player.Kick()
				return true
			}

			if len(e.InputText) < 3 {
				player.SendClientMessage("Password must be at least 3 characters long.", 1)
				registerDialog.ShowFor(player)
				return true
			}

			hashedPassword, err := argon.HashEncoded([]byte(e.InputText))

			if err != nil {
				logger.Fatal("[Player:%s] Error hashing password: %v", player.Name(), err)
				player.Kick()
				return true
			}

			user, err := q.InsertPlayer(ctx, database.InsertPlayerParams{
				Username: player.Name(),
				Password: string(hashedPassword),
			})

			if err != nil {
				logger.Fatal("[Player:%s] Error creating user: %v", player.Name(), err)
				player.Kick()
				return true
			}

			playerI := &types.PlayerI{
				Player:     player,
				StoreModel: &user,
			}
			PlayersI[playerI.ID()] = playerI
			event.Dispatch(Events, EventTypeOnAuthSuccess, &OnAuthSuccessEvent{
				PlayerI: playerI,
				Success: true,
			})
			player.SendClientMessage("Registration successful. Welcome to the server!", 1)
			player.Spawn()
			return true
		})
	} else {
		loginDialog := omp.NewPasswordDialog("Login", "Please enter your password to login.", "Login", "Cancel")
		loginDialog.ShowFor(player)

		loginDialog.On(omp.EventTypeDialogResponse, func(e *omp.InputDialogResponseEvent) bool {
			if e.Response == omp.DialogResponseRight {
				player.Kick()
				return true
			}

			verified, _ := argon2.VerifyEncoded([]byte(e.InputText), []byte(user.Password))

			if !verified {
				player.SendClientMessage("Incorrect password. Please try again.", 1)
				playerCache.LoginAttempts++
				if playerCache.LoginAttempts >= 3 {
					player.SendClientMessage("Too many failed login attempts. You've been kicked.", 1)
					time.AfterFunc(time.Millisecond*200, func() {
						player.Kick()
					})
					return true
				}
				loginDialog.ShowFor(player)
				return true
			}
			playerI := &types.PlayerI{
				Player:     player,
				StoreModel: &user,
			}
			PlayersI[playerI.ID()] = playerI
			event.Dispatch(Events, EventTypeOnAuthSuccess, &OnAuthSuccessEvent{
				PlayerI: playerI,
				Success: true,
			})
			player.SendClientMessage("Login successful. Welcome back!", 1)
			player.Spawn()
			return true
		})
	}
	return true
}

func onPlayerDisconnect(e *omp.PlayerDisconnectEvent) bool {
	player := PlayersI[e.Player.ID()]
	ctx := context.Background()
	playerPosition := player.Position()
	player.StoreModel.PosX = playerPosition.X
	player.StoreModel.PosY = playerPosition.Y
	player.StoreModel.PosZ = playerPosition.Z
	player.StoreModel.PosAngle = player.FacingAngle()
	q := database.New(database.DB)
	_, err := q.UpdatePlayer(ctx, database.UpdatePlayerParams{
		ID:       player.StoreModel.ID,
		Username: player.Name(),
		Password: player.StoreModel.Password,
		Money:    player.StoreModel.Money,
		Level:    player.StoreModel.Level,
		Exp:      player.StoreModel.Exp,
		Gold:     player.StoreModel.Gold,
		Token:    player.StoreModel.Token,
		Hour:     player.StoreModel.Hour,
		Minute:   player.StoreModel.Minute,
		Second:   player.StoreModel.Second,
		Vip:      player.StoreModel.Vip,
		Helper:   player.StoreModel.Helper,
		IsOnline: false,
		Kills:    player.StoreModel.Kills,
		Deaths:   player.StoreModel.Deaths,
		PosX:     player.StoreModel.PosX,
		PosY:     player.StoreModel.PosY,
		PosZ:     player.StoreModel.PosZ,
		PosAngle: player.StoreModel.PosAngle,
		Language: player.StoreModel.Language,
	})
	delete(PlayersI, e.Player.ID())
	if err != nil {
		logger.Fatal("[Player:%s] Error updating player: %v", e.Player.Name(), err)
		return true
	}
	return true
}

func onPlayerRequestClass(e *omp.PlayerRequestClassEvent) bool {
	e.Player.Spawn()
	return true
}

func onPlayerSpawn(e *omp.PlayerSpawnEvent) bool {
	player := PlayersI[e.Player.ID()]
	player.SetPosition(omp.Vector3{X: player.StoreModel.PosX, Y: player.StoreModel.PosY, Z: player.StoreModel.PosZ})
	player.SetFacingAngle(player.StoreModel.PosAngle)
	return true
}

func init() {
	omp.Events.Listen(omp.EventTypeGameModeInit, onGameModeInit)
	omp.Events.Listen(omp.EventTypePlayerSpawn, onPlayerSpawn)
	omp.Events.Listen(omp.EventTypePlayerRequestClass, onPlayerRequestClass)
	omp.Events.Listen(omp.EventTypePlayerConnect, onPlayerConnect)
	omp.Events.Listen(omp.EventTypePlayerDisconnect, onPlayerDisconnect)
	logger.Info("Handshake Module Initialized")
}
