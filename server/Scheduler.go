package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/mcfly722/formulator/combinator"
	"github.com/mcfly722/formulator/zeroOneTwoTree"
)

const workingTaskThresholdSec = 5
const tasksBatchSize = 5

// Scheduler ...
type Scheduler struct {
	counter           uint64
	lastSequence      string
	tasks             []*WorkingTask
	Points            *[]combinator.Point
	BestSolution      string
	SolutionDeviation float64
	ready             sync.Mutex
}

// NewScheduler ...
func NewScheduler(points *[]combinator.Point) *Scheduler {
	scheduler := &Scheduler{
		counter:           0,
		lastSequence:      "()",
		tasks:             []*WorkingTask{},
		Points:            points,
		BestSolution:      "none",
		SolutionDeviation: 10000000,
	}

	go func() {
		for {
			scheduler.ready.Lock()
			fmt.Printf(fmt.Sprintf("."))
			scheduler.ready.Unlock()
			time.Sleep(1 * time.Second)
		}
	}()

	return scheduler
}

// ToHTML ...
func (scheduler *Scheduler) ToHTML() string {
	scheduler.ready.Lock()
	out := fmt.Sprintf("loaded points: %v<br><br>best solution: %v<br><br>deviation: %v<br><br><br>", len(*scheduler.Points), scheduler.BestSolution, scheduler.SolutionDeviation)
	out += "<table border=1px cellpadding='10' cellspacing='0'><tr><td>#</td><td>Starting Sequence</td><td>Ending Sequence</td><td>Batch Size</td><td>agent</td><td>Started At</td><td>Restarted</td><td>Elapsed(sec)</td></tr>"
	for _, task := range scheduler.tasks {
		out += task.ToHTML()
	}
	out += "</table>"
	scheduler.ready.Unlock()
	return out
}

func (scheduler *Scheduler) getNewTask(agent string) (string, error) {
	scheduler.ready.Lock()

	var task *WorkingTask

	// try to search outdated task and restart it
	for i := 0; i < len(scheduler.tasks); i++ {
		if time.Now().Sub(scheduler.tasks[i].StartedAt).Seconds() > workingTaskThresholdSec {
			task = scheduler.tasks[i]
			task.Restarted++
			break
		}
	}

	// add new task
	if task == nil {

		task = &WorkingTask{
			Number:            scheduler.counter,
			StartingSequence:  scheduler.lastSequence,
			Agent:             agent,
			NumberOfSequences: tasksBatchSize,
		}

		for i := 0; i < tasksBatchSize; i++ {
			nextSequence, err := zeroOneTwoTree.GetNextBracketsSequence(scheduler.lastSequence, 2)

			if err != nil {
				panic(err)
			}

			if i < tasksBatchSize-1 {
				task.EndingSequence = nextSequence
			}

			scheduler.lastSequence = nextSequence
		}

		scheduler.counter++

		scheduler.tasks = append(scheduler.tasks, task)
	}

	task.StartedAt = time.Now()

	scheduler.ready.Unlock()
	taskString, err := json.Marshal(*task)

	if err != nil {
		return "", err
	}

	return string(taskString), nil
}
