package recursionDemo

import "fmt"

/**递归*/

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func Run() {
	fmt.Println(fact(6))
}
