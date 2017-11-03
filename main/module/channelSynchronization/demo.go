package channelSynchronization

import (
	"fmt"
	"time"
)

/**通道同步锁*/

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func Run() {
	done := make(chan bool)
	//fmt.Println(<-done) // 此处接收done消息，报：fatal error: all goroutines are asleep - deadlock!
	go worker(done)

	fmt.Println(<-done) // 接收done消息，如果不接收，worker线程无法开启
}
