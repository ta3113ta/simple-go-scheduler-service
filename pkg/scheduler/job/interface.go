package job

import "github.com/jasonlvhit/gocron"

// Job interface
type Job interface {
	// SetSchedule the time to run the job, if not set, the job is never executed
	SetSchedule(s *gocron.Scheduler)

	// Handler the job function to execute
	Handler()
}
