package rangeOverChannel

import "fmt"

/**迭代通道信息*/
func Run() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	// 通道关闭后再迭代
	for element := range queue {
		fmt.Println(element)
	}
}
