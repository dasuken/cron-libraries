package main

import (
	"fmt"
	"github.com/carlescere/scheduler"
	"runtime"
	"time"
)

func main() {
	job := func() {
		t := time.Now()
		fmt.Println("Time's up! @", t.UTC())
	}
	// Run every 2 seconds but not now.
	scheduler.Every(2).Seconds().NotImmediately().Run(job)

	// Run now and every X.
	scheduler.Every(5).Minutes().Run(job)
	scheduler.Every().Day().Run(job)
	scheduler.Every().Monday().At("08:30").Run(job)

	// Keep the program from not exiting.
	runtime.Goexit()
}