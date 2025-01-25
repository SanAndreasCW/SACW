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
	ID     uint32
	Name   string
	Payout uint32
}

type PlayerJob struct {
	Job    *Job
	OnDuty bool
}
