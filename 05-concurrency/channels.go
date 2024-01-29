package concurrency

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Job int

func longCalculation(i Job) int {
	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(duration)
	fmt.Printf("Job %d complete in %v\n", i, duration)
	return int(i) * 30
}

func makeJobs() []Job {
	jobs := make([]Job, 0, 100)
	for i := 0; i < 100; i++ {
		jobs = append(jobs, Job(rand.Intn(10000)))
	}
	return jobs
}

func Channels() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	resChannel := make(chan int)
	doneChannel := make(chan bool)
	defer close(resChannel)
	defer close(doneChannel)

	finalResult := 0
	var completedJobs int32 = 0

	jobs := makeJobs()
	for _, job := range jobs {
		go func(job Job) {
			res := longCalculation(job)
			resChannel <- res
			if int(atomic.AddInt32(&completedJobs, 1)) == len(jobs) {
				doneChannel <- true
			}
		}(job)
	}

	for {
		select {
		case res := <-resChannel:
			finalResult += res
		case <-doneChannel:
			fmt.Printf("The sum is %v\n", finalResult)
			return
		}
	}

}
