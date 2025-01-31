package commons

import (
	"github.com/SanAndreasCW/SACW/mod/enums"
	"github.com/kodeyeen/omp"
)

type Score int32

type Worker interface {
	JoinJob(job enums.JobType) bool
	LeaveJob() (Job, bool)
	SetJobScore(score Score) bool
	GiveJobScore(score Score) bool
}

type Job struct {
	ID                  enums.JobType
	Name                string
	Payout              uint32
	VehicleModels       []omp.VehicleModel
	CheckpointLocations []*omp.Vector3
	LookupLocations     []*omp.Vector3
	Group               enums.JobGroup
}

type JobCargo struct {
	Name   string
	Value  uint32
	Amount uint32
	Loaded bool
}

type PlayerJob struct {
	Job     *Job
	Company *CompanyI
	Cargo   *JobCargo
	Vehicle *omp.Vehicle
	OnDuty  bool
	Idle    bool
}
