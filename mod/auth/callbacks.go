package auth

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/colors"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/timer"
	"github.com/kodeyeen/omp"
	"github.com/matthewhartstonge/argon2"
	"time"
)

func OnGameModeInit(ctx context.Context, _ omp.Event) error {
	_, _ = omp.NewClass(255, 12, omp.Vector3{X: 0.0, Y: 0.0, Z: 0.0}, 0.0, 0, 0, 0, 0, 0, 0)
	timer.SetTimer(&timer.Timer{
		Duration: time.Minute * 1,
		Callback: func() {
			for _, playerI := range commons.PlayersI {
				if !playerI.Cache.IsLoggedIn {
					continue
				}
				playerI.StoreModel.Minute += 1

				if playerI.StoreModel.Minute >= 60 {
					playerI.StoreModel.Hour += 1
					playerI.StoreModel.Minute = 0
				}

				if companyMembership := playerI.GetCurrentCompanyMembership(ctx); companyMembership != nil {
					companyMembership.CompanyMemberInfo.Minute += 1

					if companyMembership.CompanyMemberInfo.Minute >= 60 {
						companyMembership.CompanyMemberInfo.Hour += 1
						companyMembership.CompanyMemberInfo.Minute = 0
					}
				}
			}
		},
		Async: true,
	})
	logger.Info("Auth module initialized")
	return nil
}

func onPlayerConnect(ctx context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerConnectEvent)
	argon := argon2.DefaultConfig()
	playerCache := &commons.PlayerCache{LoginAttempts: 0, IsLoggedIn: false}
	playerI := &commons.PlayerI{
		Player:      ep.Player,
		IconCounter: 0,
		Cache:       playerCache,
	}
	q := database.New(database.DB)
	user, err := q.GetPlayerByUsername(ctx, playerI.Name())

	if err != nil {
		registerDialog := omp.NewInputDialog("Registration", "Please enter your password to register.", "Register", "Cancel")
		registerDialog.ShowFor(playerI.Player)
		registerDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
			dep := e.Payload().(*omp.InputDialogResponseEvent)
			if dep.Response == omp.DialogResponseRight {
				playerI.Kick()
				return nil
			}
			if len(dep.InputText) < 3 {
				playerI.SendClientMessage("Password must be at least 3 characters long.", colors.ErrorColor.Hex)
				registerDialog.ShowFor(playerI.Player)
				return nil
			}
			hashedPassword, err := argon.HashEncoded([]byte(dep.InputText))
			if err != nil {
				logger.Fatal("[Player:%s] Error hashing password: %v", playerI.Name(), err)
				playerI.Kick()
				return nil
			}
			insertedUser, err := q.InsertPlayer(ctx, database.InsertPlayerParams{
				Username: playerI.Name(),
				Password: string(hashedPassword),
			})
			if err != nil {
				logger.Fatal("[Player:%s] Error creating user: %v", playerI.Name(), err)
				playerI.Kick()
				return nil
			}
			playerI.StoreModel = &insertedUser
			playerI.Cache.IsLoggedIn = true
			commons.PlayersI[playerI.ID()] = playerI
			onAuthEvent := omp.NewEvent(EventTypeOnAuthSuccess, &OnAuthSuccessEvent{
				PlayerI: playerI,
				Success: true,
			})
			_ = omp.EventListener().HandleEvent(ctx, onAuthEvent)
			playerI.SetMoney(playerI.StoreModel.Money)
			playerI.SendClientMessage("Registration successful. Welcome to the server!", colors.SuccessColor.Hex)
			playerI.Spawn()
			return nil
		})
	} else {
		loginDialog := omp.NewPasswordDialog("Login", "Please enter your password to login.", "Login", "Cancel")
		loginDialog.ShowFor(playerI.Player)
		loginDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
			dep := e.Payload().(*omp.InputDialogResponseEvent)
			if dep.Response == omp.DialogResponseRight {
				playerI.Kick()
				return nil
			}
			verified, _ := argon2.VerifyEncoded([]byte(dep.InputText), []byte(user.Password))
			if !verified {
				playerI.SendClientMessage("Incorrect password. Please try again.", colors.ErrorColor.Hex)
				playerCache.LoginAttempts++
				if playerCache.LoginAttempts >= 3 {
					playerI.SendClientMessage("Too many failed login attempts. You've been kicked.", colors.ErrorColor.Hex)
					time.AfterFunc(time.Millisecond*200, func() {
						playerI.Kick()
					})
					return nil
				}
				loginDialog.ShowFor(playerI.Player)
				return nil
			}
			playerI.StoreModel = &user
			playerCache.IsLoggedIn = true
			commons.PlayersI[playerI.ID()] = playerI
			onAuthEvent := omp.NewEvent(EventTypeOnAuthSuccess, &OnAuthSuccessEvent{
				PlayerI: playerI,
				Success: true,
			})
			_ = omp.EventListener().HandleEvent(ctx, onAuthEvent)
			playerI.SetMoney(playerI.StoreModel.Money)
			playerI.SendClientMessage("Login successful. Welcome back!", colors.SuccessColor.Hex)
			playerI.Spawn()
			return nil
		})
	}
	return nil
}

func onPlayerDisconnect(ctx context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerDisconnectEvent)
	playerI := commons.PlayersI[ep.Player.ID()]
	if playerI == nil || playerI.Cache == nil || !playerI.Cache.IsLoggedIn {
		return nil
	}
	playerI.SyncPlayer(ctx)
	playerI.SyncCompanyMemberInfo(ctx)
	delete(commons.PlayersI, ep.Player.ID())
	return nil
}

func onPlayerRequestClass(_ context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerRequestClassEvent)
	ep.Player.Spawn()
	return nil
}

func onPlayerSpawn(_ context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerSpawnEvent)
	playerI := commons.PlayersI[ep.Player.ID()]
	if playerI == nil || playerI.Cache == nil || !playerI.Cache.IsLoggedIn {
		ep.Player.Kick()
		return nil
	}
	playerI.SetPosition(omp.Vector3{X: playerI.StoreModel.PosX, Y: playerI.StoreModel.PosY, Z: playerI.StoreModel.PosZ})
	playerI.SetFacingAngle(playerI.StoreModel.PosAngle)
	return nil
}
