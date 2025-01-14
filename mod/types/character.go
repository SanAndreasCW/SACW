package types

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/database"
)

type PlayerI struct {
	*omp.Player
	StoreModel *database.Player
	Company    *CompanyI
}

type PlayerCache struct {
	*omp.Player
	LoginAttempts int
}

func (p *PlayerI) GetCurrentCompany() *CompanyI {
	ctx := context.Background()
	q := database.New(database.DB)
	if p.Company == nil {
		company, err := q.GetUserActiveCompany(ctx, p.StoreModel.ID)
		if err != nil {
			return nil
		}

		p.Company = &CompanyI{
			StoreModel: &company,
		}
	}
	return p.Company
}
