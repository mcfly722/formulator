package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mcfly722/formulator/combinator"

	"github.com/gorilla/mux"
)

// APIServer ...
type APIServer struct {
	router    *mux.Router
	scheduler *Scheduler
}

// NewAPIServer ...
func NewAPIServer(points *[]combinator.Point) *APIServer {
	return &APIServer{
		router:    mux.NewRouter(),
		scheduler: NewScheduler(points),
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
		out := fmt.Sprintf("<html>%v</html>", s.scheduler.ToHTML())
		io.WriteString(w, out)
	}
}

func (s *APIServer) getTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		taskString, err := s.scheduler.getNewTask(r.RemoteAddr)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		io.WriteString(w, taskString)
	}

}

func (s *APIServer) getPoints() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		pointsString, _ := json.Marshal(s.scheduler.Points)
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
