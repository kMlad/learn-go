package concurrency

import (
	"fmt"
	"math/rand"
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

	finalResult := 0
	resChannel := make(chan int)
	jobs := makeJobs()

	for _, job := range jobs {
		go func(job Job) {
			res := longCalculation(job)
			resChannel <- res
		}(job)
	}

	go func() {

		close(resChannel)

	}()

	for result := range resChannel {
		finalResult += result
	}

	fmt.Printf("The sum is %v\n", finalResult)
}
