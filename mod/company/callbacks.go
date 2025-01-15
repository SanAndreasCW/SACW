package company

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/commons"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
)

func onAuthSuccess(e *auth.OnAuthSuccessEvent) {
	if !e.Success {
		return
	}
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
		logger.Info("[Company]: Loaded %d applications for company %s", len(commons.Companies[company.ID].Applications), company.Name)
	}
	logger.Info("[Company]: Loaded %d companies", len(companies))
	return true
}
