package channelBuffering

import "fmt"

/**通道缓冲区*/
func Run() {
	message := make(chan string, 2)

	message <- "buffer1"
	message <- "buffer2"

	fmt.Println(<-message)
	fmt.Println(<-message)
}
