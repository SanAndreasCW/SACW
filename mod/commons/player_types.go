package commons

import (
	"context"
	"github.com/SanAndreasCW/SACW/mod/database"
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/SanAndreasCW/SACW/mod/logger"
	"github.com/kodeyeen/omp"
	"math/rand"
	"slices"
	"sync"
	"sync/atomic"
)

type PlayerMembership struct {
	Company           *CompanyI
	CompanyMember     *database.CompanyMember
	CompanyMemberInfo *database.CompanyMemberInfo
}

type PlayerI struct {
	Worker
	*omp.Player
	StoreModel       *database.Player
	Membership       *PlayerMembership
	CompaniesHistory []*CompanyMemberInfoI
	MoneyLock        sync.RWMutex
	IconCounter      int32
	Cache            *PlayerCache
	Job              *PlayerJob
	JobsStats        *[]database.PlayerJob
}

type PlayerCache struct {
	LoginAttempts int
	IsLoggedIn    bool
}

func (p *PlayerI) SetJobScore(score Score) {
	defer p.Job.ScoreLock.Unlock()
	p.Job.ScoreLock.Lock()
	p.Job.Score = score
}

func (p *PlayerI) GiveJobScore(score Score) {
	defer p.Job.ScoreLock.Unlock()
	p.Job.ScoreLock.Lock()
	p.Job.Score += score
}

func (p *PlayerI) JoinJob(job enums.JobType, company *CompanyI) {
	currentJob := Jobs[job]
	jobIndex := slices.IndexFunc(*p.JobsStats, func(e database.PlayerJob) bool {
		return enums.JobType(e.JobID) == currentJob.ID
	})
	js := *p.JobsStats
	jobHistory := js[jobIndex]
	p.Job = &PlayerJob{
		Job:     currentJob,
		Company: company,
		OnDuty:  true,
		Idle:    true,
		Cargo: &JobCargo{
			Name:   "Glass",
			Amount: 1,
			Loaded: false,
			Value:  1,
		},
		Score: Score(If(jobIndex == -1, 0, jobHistory.Score)),
	}
	cp := p.DefaultCheckpoint()
	cp.SetPosition(*currentJob.CheckpointLocations[rand.Intn(len(currentJob.CheckpointLocations))])
	cp.SetRadius(5.0)
}

func (p *PlayerI) LeaveJob() Job {
	job := p.Job
	p.Job = nil
	return *job.Job
}

func (p *PlayerI) SetJobCheckpoint() {
	playerVehicle, _ := p.Vehicle()
	if !slices.Contains(p.Job.Job.VehicleModels, playerVehicle.Model()) == true {
		return
	}
	p.Job.Idle = false
	pos := p.Position()
	cp := p.DefaultCheckpoint()
	if !p.Job.Cargo.Loaded {
		point := If(
			p.Job.Company == nil,
			*FindNearestToPoints(&pos, p.Job.Job.LookupLocations),
			*FindNearestToPoints(&pos, func() []*omp.Vector3 {
				var checkpoints []*omp.Vector3
				for _, checkpoint := range p.Job.Company.CompanyJobsCheckpoints[p.Job.Job.ID] {
					checkpoints = append(checkpoints, &checkpoint.CheckpointVector)
				}
				return checkpoints
			}()),
		)
		cp.SetPosition(point)
		cp.SetRadius(5.0)
		cp.Enable()
		return
	} else {
		cp.SetPosition(*FindFarthestToPoints(&pos, p.Job.Job.CheckpointLocations))
		cp.SetRadius(5.0)
		cp.Enable()
	}
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

func (p *PlayerI) GetCurrentCompanyMembership(ctx context.Context) *PlayerMembership {
	q := database.New(database.DB)
	if p.Membership == nil {
		c, err := q.GetUserActiveCompany(ctx, p.StoreModel.ID)
		if err != nil {
			return nil
		}
		p.Membership = &PlayerMembership{
			Company:           Companies[c.Company.ID],
			CompanyMember:     &c.CompanyMember,
			CompanyMemberInfo: &c.CompanyMemberInfo,
		}
	}
	return p.Membership
}

func (p *PlayerI) IsInCompany() bool {
	return If(p.Membership == nil, false, true)
}

func (p *PlayerI) HasCompanyPermission(permissions *[]int16, role int16) bool {
	return If(slices.Contains(*permissions, role), true, false)
}

func (p *PlayerI) GetPlayerCompanies(ctx context.Context) []*CompanyI {
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

func (p *PlayerI) SyncPlayer(ctx context.Context) {
	playerPosition := p.Position()
	p.StoreModel.PosX = playerPosition.X
	p.StoreModel.PosY = playerPosition.Y
	p.StoreModel.PosZ = playerPosition.Z
	p.StoreModel.PosAngle = p.FacingAngle()
	q := database.New(database.DB)
	_, err := q.UpdatePlayer(ctx, database.UpdatePlayerParams{
		ID:       p.StoreModel.ID,
		Username: p.Name(),
		Password: p.StoreModel.Password,
		Money:    p.StoreModel.Money,
		Level:    p.StoreModel.Level,
		Exp:      p.StoreModel.Exp,
		Gold:     p.StoreModel.Gold,
		Token:    p.StoreModel.Token,
		Hour:     p.StoreModel.Hour,
		Minute:   p.StoreModel.Minute,
		Vip:      p.StoreModel.Vip,
		Helper:   p.StoreModel.Helper,
		Kills:    p.StoreModel.Kills,
		Deaths:   p.StoreModel.Deaths,
		PosX:     p.StoreModel.PosX,
		PosY:     p.StoreModel.PosY,
		PosZ:     p.StoreModel.PosZ,
		PosAngle: p.StoreModel.PosAngle,
		Language: p.StoreModel.Language,
	})
	if err != nil {
		logger.Fatal("[Player:%s] Error updating: %v", p.Name(), err)
	}
}

func (p *PlayerI) SyncCompanyMemberInfo(ctx context.Context) {
	q := database.New(database.DB)
	if p.Membership != nil {
		err := q.UpdateCompanyMemberInfo(ctx, database.UpdateCompanyMemberInfoParams{
			CompanyID: p.Membership.Company.StoreModel.ID,
			PlayerID:  p.StoreModel.ID,
			Level:     p.Membership.CompanyMemberInfo.Level,
			Hour:      p.Membership.CompanyMemberInfo.Hour,
			Minute:    p.Membership.CompanyMemberInfo.Minute,
			Score:     p.Membership.CompanyMemberInfo.Score,
		})
		if err != nil {
			logger.Fatal("[Player:%s] Error updating company member info: %v", p.Name(), err)
		}
	}
}

func (p *PlayerI) SyncJobInfo(ctx context.Context) {
	q := database.New(database.DB)
	if p.Job != nil {
		_, _ = q.UpdateUserJobs(ctx, database.UpdateUserJobsParams{
			Pid: p.StoreModel.ID,
			Jid: int32(p.Job.Job.ID),
			Sc:  int32(p.Job.Score),
		})
	}
}
