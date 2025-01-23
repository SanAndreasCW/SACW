package commons

import (
	"context"
	"github.com/RahRow/omp"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"slices"
	"sync"
	"sync/atomic"
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
	IconCounter       int32
}

type PlayerCache struct {
	*omp.Player
	LoginAttempts int
	IsLoggedIn    bool
}

func (p *PlayerI) IsInCircle(centerX, centerY, radius float32) bool {
	playerPos := p.Position()
	dx := playerPos.X - centerX
	dy := playerPos.Y - centerY
	distanceSquared := (dx * dx) + (dy * dy)
	return distanceSquared <= radius*radius
}

func (p *PlayerI) NextCounter() int32 {
	c := p.IconCounter
	atomic.AddInt32(&p.IconCounter, 1)
	return c
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

func (p *PlayerI) IsInCompany() bool {
	return If(p.CompanyMemberInfo == nil, false, true)
}

func (p *PlayerI) HasCompanyPermission(permissions *[]int16, role int16) bool {
	return If(slices.Contains(*permissions, role), true, false)
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
