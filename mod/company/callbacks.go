package company

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/timer"
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
	timer.SetTimer(&timer.Timer{
		Duration: time.Duration(1) * time.Minute,
		Callback: func() {
			for _, company := range companies {
				companyI := &commons.CompanyI{
					StoreModel: &company,
				}
				commons.Companies[company.ID] = companyI
				go companyI.GiveBalance(1000)
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
