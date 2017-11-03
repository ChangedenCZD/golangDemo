package selectTimeout

import (
	"time"
	"fmt"
)

/**选择器超时*/

func Run() {
	c1 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "one"
	}()

	select {
	case msg := <-c1:
		fmt.Println("received:", msg)
	case <-time.After(time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	select {
	case msg := <-c2:
		fmt.Println("received:", msg)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
