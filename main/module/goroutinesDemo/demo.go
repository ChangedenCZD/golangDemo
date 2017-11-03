package goroutinesDemo

import "fmt"

/**线程与协程(轻量级线程)*/
func f(from string) {
	for i := 0; i < 3; i++ {
		go fmt.Println(from, ":", i)
	}
}
func Run() {
	go f("direct")

	go f("goroutine")

	go func(msg string) {
		go fmt.Println(msg)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
