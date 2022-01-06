package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mcfly722/formulator/combinator"
)

func getHTTPJSON(uri string, toObj interface{}) error {
	response, err := http.Get(uri)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, toObj)
}

// GetPoints ...
func GetPoints(serverAddr string) (*([]combinator.Point), error) {

	points := []combinator.Point{}

	err := getHTTPJSON(fmt.Sprintf("%v/getPoints", serverAddr), &points)
	if err != nil {
		return nil, err
	}

	return &points, nil
}

// Task ...
type Task struct {
	Number            uint64
	StartingSequence  string
	EndingSequence    string
	NumberOfSequences int
}

// GetTask ...
func GetTask(serverAddr string) (*Task, error) {
	var task Task
	err := getHTTPJSON(fmt.Sprintf("%v/getTask", serverAddr), &task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
