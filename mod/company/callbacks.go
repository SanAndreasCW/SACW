package company

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/timer"
	"github.com/kodeyeen/omp"
	"math/rand"
	"time"
)

func onAuthSuccess(_ context.Context, e omp.Event) error {
	ep := e.Payload().(*auth.OnAuthSuccessEvent)
	if !ep.Success {
		return nil
	}
	playerI := ep.PlayerI
	go func() {
		for _, company := range commons.Companies {
			playerI.SetMapIcon(int(playerI.NextCounter()), 52, 0, enums.MapiconLocal, omp.Vector3{
				X: company.CompanyOffice.IconX,
				Y: company.CompanyOffice.IconY,
				Z: 0.0,
			})
		}
	}()
	companyMembership := ep.PlayerI.GetCurrentCompanyMembership()
	if companyMembership != nil {
		companyMembership.Company.ReloadApplications()
	}
	return nil
}

func OnGameModeInit(ctx context.Context, _ omp.Event) error {
	q := database.New(database.DB)
	companies, err := q.GetCompanies(ctx)
	if err != nil {
		logger.Fatal("[Company]: Failed to load companies: %v", err)
		return nil
	}
	for _, company := range companies {
		pickup, _ := omp.NewPickup(1239, 1, -1, omp.Vector3{
			X: company.CompanyOffice.PickupX,
			Y: company.CompanyOffice.PickupY,
			Z: company.CompanyOffice.PickupZ,
		})
		companyJobs, _ := q.GetCompanyJobs(ctx, company.Company.ID)
		companyI := &commons.CompanyI{
			StoreModel:    &company.Company,
			CompanyOffice: &company.CompanyOffice,
			CompanyPickup: pickup,
			CompanyJobs:   companyJobs,
		}
		commons.Companies[company.Company.ID] = companyI
	}
	timer.SetTimer(&timer.Timer{
		Duration: time.Duration(30) * time.Minute,
		Callback: func() {
			for _, companyI := range commons.Companies {
				multiplierSign := commons.If[float32](rand.Float32() > 0.5, 1.0, -1.0)
				companyI.StoreModel.Multiplier += companyI.StoreModel.Multiplier + (multiplierSign * rand.Float32() / 10)
			}
		},
	})
	timer.SetTimer(&timer.Timer{
		Duration: time.Duration(1) * time.Minute,
		Callback: func() {
			for _, companyI := range commons.Companies {
				go companyI.GiveBalance(int32(1000.0 * companyI.StoreModel.Multiplier))
				go companyI.ReloadApplications()
			}
		},
	})
	logger.Info("[Company]: Loaded %d companies", len(companies))
	return nil
}

func OnGameModeExit(ctx context.Context, _ omp.Event) error {
	q := database.New(database.DB)
	for _, company := range commons.Companies {
		err := q.UpdateCompany(ctx, database.UpdateCompanyParams{
			ID:         company.StoreModel.ID,
			Multiplier: company.StoreModel.Multiplier,
			Balance:    company.StoreModel.Balance,
		})

		if err != nil {
			logger.Fatal("[Company]: Failed to update specific company %s: %v", company.StoreModel.Name, err)
		}
	}
	return nil
}

func onPlayerKeyStateChange(_ context.Context, e omp.Event) error {
	ep := e.Payload().(*omp.PlayerKeyStateChangeEvent)
	playerI := commons.PlayersI[ep.Player.ID()]
	keys := playerI.KeyData()
	if keys.Keys == omp.PlayerKeyWalk {
		for _, company := range commons.Companies {
			pickupPosition := company.CompanyPickup.Position()
			if playerI.IsInCircle(pickupPosition.X, pickupPosition.Y, 5.0) {
				companyOptionSelectionDialog := omp.NewListDialog("Select Your Action", "Select", "Close")
				companyOptionSelectionDialog.Add("Stats")
				if playerI.Job == nil {
					companyOptionSelectionDialog.Add("Jobs")
				} else {
					companyOptionSelectionDialog.Add("Abandon Job")
				}
				if playerI.IsInCompany() {
					if playerI.Membership.Company == company {
						if playerI.HasCompanyPermission(
							&commons.CompanyApplicationPermissions,
							playerI.Membership.CompanyMember.Role) {
							companyOptionSelectionDialog.Add("Applications")
						}
					}
				} else {
					companyOptionSelectionDialog.Add("Send Application")
				}
				companyOptionSelectionDialog.Events.ListenFunc(omp.EventTypeDialogResponse, func(ctx context.Context, e omp.Event) error {
					ep := e.Payload().(*omp.ListDialogResponseEvent)
					if ep.Response == omp.DialogResponseRight {
						return nil
					}
					switch ep.Item {
					case "Applications":
						companyApplicationsActions(playerI)
						return nil
					case "Stats":
						statsDialog := companyStatsDialog(company)
						statsDialog.ShowFor(playerI.Player)
						return nil
					case "Send Application":
						companiesApplicationAction(playerI, &company.StoreModel.Tag)
						return nil
					case "Jobs":
						companiesJobsAction(playerI, company)
						return nil
					case "Abandon Job":
						companiesJobAbandonAction(playerI)
						return nil
					}
					return nil
				})
				companyOptionSelectionDialog.ShowFor(playerI.Player)
				return nil
			}
		}
	}
	return nil
}
