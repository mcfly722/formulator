package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mcfly722/formulator/zeroOneTwoTree"
)

var (
	serverAddrFlag    *string
	errorSleepSecFlag *int
)

func recombine(sequence string) {
	fmt.Println(fmt.Sprintf("%v", sequence))
}

func main() {
	serverAddrFlag = flag.String("server", "http://127.0.0.1:8080", "server address")
	errorSleepSecFlag = flag.Int("errorSleepSec", 3, "sleep * seconds if connection error catched")

	fmt.Println(fmt.Sprintf("working with server: %v", *serverAddrFlag))

	points, err := GetPoints(*serverAddrFlag)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("obtained %v points", len(*points)))

	for {
		task, err := GetTask(*serverAddrFlag)

		if err != nil {
			fmt.Println(fmt.Sprintf("error: %v sleeping %v sec", err, *errorSleepSecFlag))
			time.Sleep(time.Second * time.Duration(*errorSleepSecFlag))
		} else {

			fmt.Println(fmt.Sprintf("%v", task))

			sequence := task.StartingSequence
			for i := 0; i < task.NumberOfSequences; i++ {

				recombine(sequence)

				nextSequence, err := zeroOneTwoTree.GetNextBracketsSequence(sequence, 2)
				if err != nil {
					panic(err)
				}
				sequence = nextSequence
			}

			time.Sleep(time.Second)
		}
	}
}
