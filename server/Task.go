package main

import (
	"fmt"
	"time"
)

// WorkingTask ...
type WorkingTask struct {
	Number            uint64
	StartingSequence  string
	EndingSequence    string
	NumberOfSequences int
	Agent             string
	StartedAt         time.Time
	Restarted         uint64
	FinishedAt        time.Time
}

// ToHTML ...
func (task *WorkingTask) ToHTML() string {
	var duration time.Duration

	if task.FinishedAt.IsZero() {
		duration = time.Now().Sub(task.StartedAt)
	} else {
		duration = time.Now().Sub(task.FinishedAt)
	}

	return fmt.Sprintf("<tr><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td><td>%v</td></tr>", task.Number, task.StartingSequence, task.EndingSequence, tasksBatchSize, task.Agent, task.StartedAt.Format(time.RFC3339), task.Restarted, uint64(duration.Seconds()))
}
