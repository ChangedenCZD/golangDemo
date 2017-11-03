package statefulGoroutine

import (
	"math/rand"
	"sync/atomic"
	"time"
	"fmt"
)

/**带状态的线程*/

type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func Run() {
	var readOps uint64 = 0
	var writeOps uint64 = 0
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{rand.Intn(5), make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{rand.Intn(5),
					rand.Intn(100), make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond * 100)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("readOps:", atomic.LoadUint64(&readOps))
	fmt.Println("writeOps:", atomic.LoadUint64(&writeOps))

	close(reads)
}
