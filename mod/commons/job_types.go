package commons

type Score int32

type Worker interface {
}

type JobI struct {
	Worker
}

type PlayerJob struct {
	Job *JobI
}
