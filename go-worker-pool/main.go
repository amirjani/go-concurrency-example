package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for i := range jobs {
		fmt.Printf("Jobs %d started |||| ", i)
		time.Sleep(time.Second * 1)
		fmt.Printf("Jobs %d finished |||| ", i)
		results <- i * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 0; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
