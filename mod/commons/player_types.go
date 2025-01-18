package commons

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"sync"
)

type PlayerMembership struct {
	Company       *CompanyI
	CompanyMember *database.CompanyMember
}

type PlayerI struct {
	*omp.Player
	StoreModel        *database.Player
	CompanyMemberInfo *PlayerMembership
	CompaniesHistory  []*CompanyMemberInfoI
	MoneyLock         sync.RWMutex
}

type PlayerCache struct {
	*omp.Player
	LoginAttempts int
	IsLoggedIn    bool
}

func (p *PlayerI) GetCurrentCompanyMembership() *PlayerMembership {
	ctx := context.Background()
	q := database.New(database.DB)
	if p.CompanyMemberInfo == nil {
		c, err := q.GetUserActiveCompany(ctx, p.StoreModel.ID)
		if err != nil {
			logger.Fatal("[PlayerI] GetCurrentCompanyMembership Error:", err)
			return nil
		}
		p.CompanyMemberInfo = &PlayerMembership{
			Company:       Companies[c.Company.ID],
			CompanyMember: &c.CompanyMember,
		}
	}
	return p.CompanyMemberInfo
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

func (p *PlayerI) GiveMoney(money int32) {
	p.MoneyLock.Lock()
	p.Player.GiveMoney(int(money))
	p.StoreModel.Money = p.StoreModel.Money + money
	p.MoneyLock.Unlock()
}

func (p *PlayerI) SetMoney(money int32) {
	p.MoneyLock.Lock()
	p.Player.SetMoney(int(money))
	p.StoreModel.Money = money
	p.MoneyLock.Unlock()
}
