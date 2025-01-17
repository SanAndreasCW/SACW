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
				companyI.ReloadApplications()
			}
		},
	})
	logger.Info("[Company]: Loaded %d companies", len(companies))
	return true
}
