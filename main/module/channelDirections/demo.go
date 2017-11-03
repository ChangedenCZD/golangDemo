package channelDirections

import "fmt"

/**通道方向*/
func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	//<-chan 传入一个用于接收信息的通道
	//chan<- 传入一个用于输入信息的通道
	msg := <-pings
	pongs <- msg
}

func Run() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
