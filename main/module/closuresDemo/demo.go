package closuresDemo

import "fmt"

/**闭包*/

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func Run() {
	nextInt1 := intSeq()
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())
	fmt.Println(nextInt1())

	nextInt2 := intSeq()
	fmt.Println(nextInt2())
}
