package jobb

import (
	"fmt"

	"scheduler/pkg/database"

	"github.com/jasonlvhit/gocron"
)

type JobB struct {
	db *database.Database
}

func New(db *database.Database) *JobB {
	return &JobB{
		db: db,
	}
}

func (j *JobB) SetSchedule(s *gocron.Scheduler) {
	s.Every(1).Second().Do(j.Handler)

}

func (j *JobB) Handler() {
	dbPrt := j.db.DBAddress()
	fmt.Printf("JobB: %s\n", dbPrt)
}
