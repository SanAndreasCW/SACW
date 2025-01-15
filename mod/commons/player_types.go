package commons

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/database"
)

type PlayerI struct {
	*omp.Player
	StoreModel *database.Player
	Company    *CompanyI
	Companies  []*CompanyMemberInfoI
}

type PlayerCache struct {
	*omp.Player
	LoginAttempts int
}

func (p *PlayerI) GetCurrentCompany() *CompanyI {
	ctx := context.Background()
	q := database.New(database.DB)
	if p.Company == nil {
		c, err := q.GetUserActiveCompany(ctx, p.StoreModel.ID)
		if err != nil {
			return nil
		}
		p.Company = Companies[c.ID]
	}
	return p.Company
}

func (p *PlayerI) GetPlayerCompanies() []*CompanyI {
	ctx := context.Background()
	q := database.New(database.DB)
	companiesInfo, err := q.GetUserCompaniesInfo(ctx, p.StoreModel.ID)
	if err != nil {
		return nil
	}
	var companiesInfoI []*CompanyI
	for _, companyInfo := range companiesInfo {
		companiesInfoI = append(companiesInfoI, &CompanyI{
			StoreModel: &companyInfo,
		})
	}
	return companiesInfoI
}
