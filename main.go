package main

import (
	"fmt"

	"github.com/jasonlvhit/gocron"
)

type Server struct {
	scheduler *Scheduler
}

func NewServer() *Server {
	db := &Database{}

	jobs := []Job{
		&JobA{
			db: db,
		},
		&JobB{
			db: db,
		},
	}

	scheduler := &Scheduler{
		Scheduler: gocron.NewScheduler(),
		jobs:      jobs,
	}

	scheduler.Build()

	return &Server{
		scheduler: scheduler,
	}
}

func (s *Server) Start() {
	<-s.scheduler.Start()
}

type Database struct{}

func (d *Database) DBAddress() string {
	return fmt.Sprintf("%p", d)
}

// Job interface
type Job interface {
	// SetSchedule the time to run the job, if not set, the job is never executed
	SetSchedule(s *gocron.Scheduler)

	// Handler the job function to execute
	Handler()
}

type Scheduler struct {
	*gocron.Scheduler

	jobs []Job
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Scheduler: gocron.NewScheduler(),
	}
}

func (s *Scheduler) Build() {
	for _, job := range s.jobs {
		job.SetSchedule(s.Scheduler)
	}
}

type JobA struct {
	db *Database
}

func (j *JobA) SetSchedule(s *gocron.Scheduler) {
	s.Every(1).Second().Do(j.Handler)
}

func (j *JobA) Handler() {
	dbPrt := j.db.DBAddress()
	fmt.Printf("JobA: %s\n", dbPrt)
}

type JobB struct {
	db *Database
}

func (j *JobB) SetSchedule(s *gocron.Scheduler) {
	s.Every(3).Second().Do(j.Handler)
}

func (j *JobB) Handler() {
	dbPrt := j.db.DBAddress()
	fmt.Printf("JobB: %s\n", dbPrt)
}

func main() {
	server := NewServer()
	server.Start()
}
