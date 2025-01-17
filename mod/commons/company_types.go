package commons

import (
	"context"
	"database/sql"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"sync"
)

type CompanyApplicationI struct {
	StoreModel  *database.CompanyApplication
	PlayerModel *database.Player
}

type CompanyI struct {
	StoreModel        *database.Company
	Applications      []*CompanyApplicationI
	ApplicationsMutex *sync.RWMutex
	Members           []*PlayerI
	MembersLock       *sync.RWMutex
}

type CompanyMemberInfoI struct {
	StoreModel *database.CompanyMemberInfo
}

func (ci *CompanyI) ReloadApplications() {
	ctx := context.Background()
	q := database.New(database.DB)
	applications, err := q.GetCompanyApplications(ctx, database.GetCompanyApplicationsParams{
		CompanyID: ci.StoreModel.ID,
		Accepted:  enums.OnProgress,
	})
	if err != nil {
		logger.Fatal("[CompanyApplications]: Couldn't load applications: %v", err)
	}
	ci.Applications = nil
	for _, application := range applications {
		ci.Applications = append(ci.Applications, &CompanyApplicationI{
			StoreModel:  &application.CompanyApplication,
			PlayerModel: &application.Player,
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

func (ci *CompanyI) AddApplication(application *database.CompanyApplication) {
	ci.ApplicationsMutex.Lock()
	defer ci.ApplicationsMutex.Unlock()
	ci.Applications = append(ci.Applications, &CompanyApplicationI{
		StoreModel: application,
	})
}

func (ci *CompanyI) CreateApplication(playerI *PlayerI, description string) bool {
	ctx := context.Background()
	q := database.New(database.DB)
	application, err := q.InsertCompanyApplication(ctx, database.InsertCompanyApplicationParams{
		PlayerID:    playerI.StoreModel.ID,
		CompanyID:   ci.StoreModel.ID,
		Description: sql.NullString{String: description, Valid: true},
	})
	if err != nil {
		logger.Fatal("[CreateApplication]: Couldn't create application: %v", err)
		return false
	}
	ci.AddApplication(&application)
	return true
}

func (ci *CompanyI) AddMember(memberI *PlayerI) {
	ci.MembersLock.Lock()
	ci.Members = append(ci.Members, memberI)
	ci.MembersLock.Unlock()
}
