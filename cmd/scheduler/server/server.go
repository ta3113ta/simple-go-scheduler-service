package server

import (
	"time"

	"scheduler/pkg/database"
	"scheduler/pkg/scheduler"
	"scheduler/pkg/scheduler/job"
	"scheduler/pkg/scheduler/job/joba"
	"scheduler/pkg/scheduler/job/jobb"
)

type Server struct {
	scheduler *scheduler.Scheduler
}

func New() *Server {
	db := &database.Database{}

	jobs := []job.Job{
		joba.New(db),
		jobb.New(db),
	}

	scheduler := scheduler.New(jobs)
	scheduler.Build()

	return &Server{
		scheduler: scheduler,
	}
}

func (s *Server) Start() {
	s.scheduler.ChangeLoc(time.UTC)
	<-s.scheduler.Start()
}
