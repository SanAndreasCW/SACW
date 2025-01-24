package auth

import (
	"context"
	"fmt"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/timer"
	"github.com/kodeyeen/event"
	"time"

	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/matthewhartstonge/argon2"
)

func onGameModeInit(_ *omp.GameModeInitEvent) bool {
	_, _ = omp.NewClass(255, 12, omp.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, 0.0, 0, 0, 0, 0, 0, 0)
	timer.SetTimer(&timer.Timer{
		Duration: time.Minute * 1,
		Callback: func() {
			for _, playerI := range commons.PlayersI {
				playerI.StoreModel.Minute += 1

				if playerI.StoreModel.Minute >= 60 {
					playerI.StoreModel.Hour += 1
					playerI.StoreModel.Minute = 0
				}

				if companyMembership := playerI.GetCurrentCompanyMembership(); companyMembership != nil {

				}
			}
		},
		Async: true,
	})
	return true
}

func onPlayerConnect(e *omp.PlayerConnectEvent) bool {
	argon := argon2.DefaultConfig()
	player := e.Player
	playerCache := &commons.PlayerCache{Player: player, LoginAttempts: 0, IsLoggedIn: false}
	commons.PlayersCache[player.ID()] = playerCache
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
			insertedUser, err := q.InsertPlayer(ctx, database.InsertPlayerParams{
				Username: player.Name(),
				Password: string(hashedPassword),
			})
			if err != nil {
				logger.Fatal("[Player:%s] Error creating user: %v", player.Name(), err)
				player.Kick()
				return true
			}
			playerI := &commons.PlayerI{
				Player:      player,
				StoreModel:  &insertedUser,
				IconCounter: 0,
			}
			commons.PlayersI[playerI.ID()] = playerI
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
			playerI := &commons.PlayerI{
				Player:      player,
				StoreModel:  &user,
				IconCounter: 0,
			}
			playerCache.IsLoggedIn = true
			commons.PlayersI[playerI.ID()] = playerI
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
	playerCache := commons.PlayersCache[e.Player.ID()]
	if !playerCache.IsLoggedIn {
		return true
	}
	player := commons.PlayersI[e.Player.ID()]
	player.SyncPlayer()
	player.SyncCompanyMemberInfo()
	delete(commons.PlayersI, e.Player.ID())
	delete(commons.PlayersCache, e.Player.ID())
	return true
}

func onPlayerRequestClass(e *omp.PlayerRequestClassEvent) bool {
	e.Player.Spawn()
	return true
}

func onPlayerSpawn(e *omp.PlayerSpawnEvent) bool {
	playerCache := commons.PlayersCache[e.Player.ID()]
	if !playerCache.IsLoggedIn {
		e.Player.Kick()
		return true
	}
	playerI := commons.PlayersI[e.Player.ID()]
	playerI.SetPosition(omp.Vector3{X: playerI.StoreModel.PosX, Y: playerI.StoreModel.PosY, Z: playerI.StoreModel.PosZ})
	playerI.SetFacingAngle(playerI.StoreModel.PosAngle)
	return true
}

func onPlayerText(e *omp.PlayerTextEvent) bool {
	msg := fmt.Sprintf("[ID:%d|Name:%s]: %s", e.Player.ID(), e.Player.Name(), e.Message)
	commons.SendClientMessageToAll(msg, 0xFFFFFFFF)
	return true
}
