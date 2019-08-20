package manager

import (
	"fmt"
	"runtime"
	"time"

	"github.com/mendrugory/comiccon/data"
	"github.com/mendrugory/comiccon/worker"
)

func Run(workers int, d data.Resource) chan bool {

	runtime.GOMAXPROCS(workers)

	end := make(chan bool)

	go doRun(workers, d, end)

	return end
}

func doRun(workers int, d data.Resource, end chan bool) {

	listen(workers, d)

	end <- true
}

func listen(maxWorkers int, d data.Resource) {

	activeWorkers := 0
	maxTimeouts := 5
	timeoutsCount := 0
	timeoutTime := 2 // seconds
	maxJobQueue := 1000

	releaseC := make(chan bool)
	jobsC := make(chan data.Resource, maxJobQueue)

	jobsC <- d

	var waitingJobs []data.Resource

L:
	for {
		select {
		case <-releaseC:
			if len(waitingJobs) > 0 {
				launch(waitingJobs[0], releaseC, jobsC)
				waitingJobs = waitingJobs[1:]
			} else {
				activeWorkers--
			}
			timeoutsCount = 0

		case d := <-jobsC:
			if activeWorkers < maxWorkers {
				activeWorkers++
				launch(d, releaseC, jobsC)
			} else {
				waitingJobs = append(waitingJobs, d)
			}
			timeoutsCount = 0

		default:
			if activeWorkers > 0 {
				timeoutsCount = 0
			} else {
				if timeoutsCount == maxTimeouts {
					fmt.Println("No more jobs to complete.")
					break L
				}
				timeoutsCount++
				fmt.Println("Timeouts: ", timeoutsCount)
				time.Sleep(time.Duration(timeoutTime) * time.Second)
			}
		}
	}

}

func launch(d data.Resource, releaseC chan bool, jobsC chan data.Resource) {
	go worker.Download(d, releaseC, jobsC)
}
