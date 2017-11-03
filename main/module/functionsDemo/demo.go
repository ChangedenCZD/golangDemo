package functionsDemo

/**方法*/

import "fmt"

func plus1(a int, b int) int {
	return a + b
}

func plus2(a, b, c int) int {
	return plus1(plus1(a, b), c)
}

// 可变参数
func plus3(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func Run() {
	fmt.Println("1 + 2 = ", plus1(1, 2))
	fmt.Println("1 + 2 + 3 = ", plus2(1, 2, 3))
	fmt.Println("1 + 2 + 3 + ... + 9 = ", plus3(1, 2, 3, 4, 5, 6, 7, 8, 9))
}
