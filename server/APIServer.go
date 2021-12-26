package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/mcfly722/formulator/combinator"
	"github.com/mcfly722/formulator/zeroOneTwoTree"

	"github.com/gorilla/mux"
)

const tasksBatchSize = 5

// APIServer ...
type APIServer struct {
	router  *mux.Router
	context *Context
}

// Context ...
type Context struct {
	counter      uint64
	lastSequence string
	tasks        []*WorkingTask
	Points       *[]combinator.Point
	ready        sync.Mutex
}

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

const workingTaskThresholdSec = 5

func (ctx *Context) getNewTask(agent string) (string, error) {
	ctx.ready.Lock()

	var task *WorkingTask

	// try to search outdated task and restart it
	for i := 0; i < len(ctx.tasks); i++ {
		if time.Now().Sub(ctx.tasks[i].StartedAt).Seconds() > workingTaskThresholdSec {
			task = ctx.tasks[i]
			task.Restarted++
			break
		}
	}

	// add new task
	if task == nil {

		task = &WorkingTask{
			Number:            ctx.counter,
			StartingSequence:  ctx.lastSequence,
			Agent:             agent,
			NumberOfSequences: tasksBatchSize,
		}

		for i := 0; i < tasksBatchSize; i++ {
			nextSequence, err := zeroOneTwoTree.GetNextBracketsSequence(ctx.lastSequence, 2)

			if err != nil {
				panic(err)
			}

			if i < tasksBatchSize-1 {
				task.EndingSequence = nextSequence
			}

			ctx.lastSequence = nextSequence
		}

		ctx.counter++

		ctx.tasks = append(ctx.tasks, task)
	}

	task.StartedAt = time.Now()

	ctx.ready.Unlock()
	taskString, err := json.Marshal([]WorkingTask{*task})

	if err != nil {
		return "", err
	}

	return string(taskString), nil
}

// ToHTML ...
func (ctx *Context) ToHTML() string {
	ctx.ready.Lock()
	out := "<table border=1px cellpadding='10' cellspacing='0'><tr><td>#</td><td>Starting Sequence</td><td>Ending Sequence</td><td>Batch Size</td><td>agent</td><td>Started At</td><td>Restarted</td><td>Elapsed(sec)</td></tr>"
	for _, task := range ctx.tasks {
		out += task.ToHTML()
	}
	out += "</table>"
	ctx.ready.Unlock()
	return out
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

// NewAPIServer ...
func NewAPIServer(_points *[]combinator.Point) *APIServer {
	return &APIServer{
		router: mux.NewRouter(),
		context: &Context{
			counter:      0,
			lastSequence: "()",
			tasks:        []*WorkingTask{},
			Points:       _points,
		},
	}
}

// Start ...
func (s *APIServer) Start(bindAddr string) error {
	s.router.HandleFunc("/", s.handleTasks())
	s.router.HandleFunc("/getTask", s.getTask())
	s.router.HandleFunc("/getPoints", s.getPoints())

	fmt.Println(fmt.Sprintf("starting server on %v", bindAddr))
	return http.ListenAndServe(bindAddr, s.router)
}

func (s *APIServer) handleTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		out := fmt.Sprintf("<html>%v</html>", s.context.ToHTML())
		io.WriteString(w, out)
	}

}

func (s *APIServer) getTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		taskString, err := s.context.getNewTask(r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		io.WriteString(w, taskString)
	}

}

func (s *APIServer) getPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pointsString, _ := json.Marshal(s.context.Points)
		io.WriteString(w, string(pointsString))
	}

}

var (
	bindAddrFlag    *string
	stateFileFlag   *string
	samplesFileFlag *string
)

func main() {

	bindAddrFlag = flag.String("bindAddr", "127.0.0.1:8080", "bind address")
	stateFileFlag = flag.String("stateFile", "state.json", "file for current state")
	samplesFileFlag = flag.String("samplesFile", "..\\samples\\exponent\\exponent.json", "file of points for required function")

	body, err := ioutil.ReadFile(*samplesFileFlag)
	if err != nil {
		panic(err)
	}

	points := []combinator.Point{}

	err = json.Unmarshal(body, &points)
	if err != nil {
		panic(err)
	}

	server := NewAPIServer(&points)
	if err := server.Start(*bindAddrFlag); err != nil {
		log.Fatal(err)
	}

}
