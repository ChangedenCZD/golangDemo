package nonBlockingChannelOperations

import "fmt"

/**非阻塞通道的操作*/
func Run() {
	message := make(chan string)
	signal := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signal:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
