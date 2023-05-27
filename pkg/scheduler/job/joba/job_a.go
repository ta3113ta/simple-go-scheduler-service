package joba

import (
	"fmt"

	"scheduler/pkg/database"

	"github.com/jasonlvhit/gocron"
)

type JobA struct {
	db *database.Database
}

func New(db *database.Database) *JobA {
	return &JobA{
		db: db,
	}
}

func (j *JobA) SetSchedule(s *gocron.Scheduler) {
	s.Every(3).Second().Do(j.Handler)
}

func (j *JobA) Handler() {
	dbPrt := j.db.DBAddress()
	fmt.Printf("JobA: %s\n", dbPrt)
}
