package company

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/timer"
	"math/rand"
	"time"
)

func onAuthSuccess(e *auth.OnAuthSuccessEvent) bool {
	if !e.Success {
		return true
	}
	companyMembership := e.PlayerI.GetCurrentCompanyMembership()
	companyMembership.Company.ReloadApplications()
	return true
}

func onGameModeInit(_ *omp.GameModeInitEvent) bool {
	ctx := context.Background()
	q := database.New(database.DB)
	companies, err := q.GetCompanies(ctx)

	if err != nil {
		logger.Fatal("[Company]: Failed to load companies: %v", err)
		return true
	}
	for _, company := range companies {
		companyI := &commons.CompanyI{
			StoreModel:    &company.Company,
			CompanyOffice: &company.CompanyOffice,
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
	return true
}

func onGameModeExit(_ *omp.GameModeExitEvent) bool {
	ctx := context.Background()
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
	return true
}
