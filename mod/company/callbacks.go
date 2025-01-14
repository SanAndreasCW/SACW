package company

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/auth"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/SanAndreasCW/SACW/mod/types"
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
		Companies[company.ID] = &types.CompanyI{
			StoreModel: &company,
		}
	}
	logger.Info("[Company]: Loaded %d companies", len(companies))
	return true
}
