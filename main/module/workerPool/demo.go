package workerPool

import (
	"fmt"
	"time"
)

/**线程池*/

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j
	}
}

func Run() {
	startTime := time.Now().UnixNano() / 1000000
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for j := 1; j <= 5; j++ {
		jobs <- j
	}

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
	fmt.Println("耗时：", time.Now().UnixNano()/1000000-startTime)
}
