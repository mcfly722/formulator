package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// APIServer ...
type APIServer struct {
	router  *mux.Router
	context *Context
}

// Context ...
type Context struct {
	counter int
}

// NewAPIServer ...
func NewAPIServer() *APIServer {
	return &APIServer{
		router:  mux.NewRouter(),
		context: &Context{counter: 0},
	}
}

// Start ...
func (s *APIServer) Start(bindAddr string) error {
	s.router.HandleFunc("/hello", s.handleHello())

	fmt.Printf(fmt.Sprintf("starting server on %v", bindAddr))
	return http.ListenAndServe(bindAddr, s.router)
}

func (s *APIServer) handleHello() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("Hello #%v", s.context.counter))
		s.context.counter++
	}
}

var (
	bindAddrFlag  *string
	stateFileFlag *string
)

func main() {

	bindAddrFlag = flag.String("bindAddr", ":8080", "bind address")
	stateFileFlag = flag.String("stateFile", "state.json", "file for current state")

	server := NewAPIServer()
	if err := server.Start(*bindAddrFlag); err != nil {
		log.Fatal(err)
	}

}
