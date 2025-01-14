package types

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
)

type CompanyApplicationI struct {
	StoreModel *database.CompanyApplication
}

type CompanyI struct {
	StoreModel   *database.Company
	Applications []*CompanyApplicationI
	Members      []*PlayerI
}

func (ci *CompanyI) ReloadApplications() {
	ctx := context.Background()
	q := database.New(database.DB)

	applications, err := q.GetCompanyApplications(ctx, ci.StoreModel.ID)
	if err != nil {
		logger.Fatal("[CompanyApplications]: Couldn't load applications: %v", err)
	}
	ci.Applications = nil
	for _, application := range applications {
		ci.Applications = append(ci.Applications, &CompanyApplicationI{
			StoreModel: &application,
		})
	}
}
