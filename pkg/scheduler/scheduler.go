package scheduler

import (
	"scheduler/pkg/scheduler/job"

	"github.com/jasonlvhit/gocron"
)

type Scheduler struct {
	*gocron.Scheduler

	jobs []job.Job
}

func New(jobs []job.Job) *Scheduler {
	return &Scheduler{
		Scheduler: gocron.NewScheduler(),
		jobs:      jobs,
	}
}

func (s *Scheduler) Build() {
	for i, job := range s.jobs {
		job.SetSchedule(s.Scheduler)
		s.checkSetupSchedule(i)
	}
}

func (s *Scheduler) checkSetupSchedule(idx int) {
	// if the job schedule is set correctly, the length of the scheduler should be idx + 1
	if s.Scheduler.Len()-1 != idx {
		panic("job schedule is not set correctly, please set the schedule in SetSchedule method")
	}
}
