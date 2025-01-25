package commons

import "github.com/SanAndreasCW/SACW/mod/enums"

type Score int32

type Worker interface {
	Join(job enums.JobType) bool
	Leave() (Job, bool)
	SetScore(score Score) bool
	GiveScore(score Score) bool
}

type Job struct {
	ID     enums.JobType
	Name   string
	Payout uint32
}

type PlayerJob struct {
	Job     *Job
	Company *CompanyI
	OnDuty  bool
}
