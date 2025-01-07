package company

import (
	"context"
	"github.com/LosantosGW/go_LSGW/mod/auth"
	"github.com/LosantosGW/go_LSGW/mod/database"
	"github.com/LosantosGW/go_LSGW/mod/logger"
	"github.com/LosantosGW/go_LSGW/mod/types"
	"github.com/RahRow/omp"
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

	for index, company := range companies {
		Companies[index] = &types.CompanyI{
			StoreModel: &company,
		}
	}

	logger.Info("[Company]: Loaded %d companies", len(companies))
	return true
}
