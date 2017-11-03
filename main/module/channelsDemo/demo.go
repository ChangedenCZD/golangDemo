package channelsDemo

import "fmt"

/**通道*/
func Run() {
	message := make(chan string)
	go func() {
		message <- "ping"
	}()

	go func() {
		message <- "pong"
	}()

	msg := <-message
	fmt.Println(msg)

	msg = <-message
	fmt.Println(msg)
}
