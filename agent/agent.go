package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	serverAddrFlag    *string
	errorSleepSecFlag *int
)

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
			time.Sleep(time.Second)
		}
	}
}
