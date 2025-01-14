package types

import (
	"context"
	"database/sql"
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

type CompanyMemberInfoI struct {
	StoreModel *database.CompanyMemberInfo
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

func (ci *CompanyI) AnswerApplication(playerI *PlayerI, answer int16) bool {
	ctx := context.Background()
	q := database.New(database.DB)
	err := q.AnswerCompanyApplication(ctx, database.AnswerCompanyApplicationParams{
		PlayerID:  playerI.StoreModel.ID,
		CompanyID: ci.StoreModel.ID,
		Accepted:  answer,
	})
	if err != nil {
		return false
	}
	return true
}

func (ci *CompanyI) CreateApplication(playerI *PlayerI, description string) bool {
	ctx := context.Background()
	q := database.New(database.DB)
	application, err := q.InsertCompanyApplication(ctx, database.InsertCompanyApplicationParams{
		PlayerID:    playerI.StoreModel.ID,
		CompanyID:   ci.StoreModel.ID,
		Description: sql.NullString{String: description, Valid: true},
	})
	ci.Applications = append(ci.Applications, &CompanyApplicationI{
		StoreModel: &application,
	})
	if err != nil {
		return false
	} else {
		return true
	}
}
