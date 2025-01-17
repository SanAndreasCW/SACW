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

func onAuthSuccess(e *auth.OnAuthSuccessEvent) {
	if !e.Success {
		return
	}
	companyMembership := e.PlayerI.GetCurrentCompanyMembership()
	companyMembership.Company.ReloadApplications()
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
		commons.Companies[company.ID] = &commons.CompanyI{
			StoreModel: &company,
		}
		commons.Companies[company.ID].ReloadApplications()
		timer.SetTimer(&timer.Timer{
			Duration: time.Duration(1) * time.Minute,
			Callback: commons.Companies[company.ID].ReloadApplications,
		})
	}
	logger.Info("[Company]: Loaded %d companies", len(companies))
	return true
}
