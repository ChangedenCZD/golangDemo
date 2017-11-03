package sliceDemo
/**数组分割*/
import (
	"fmt"
)

func Run() {
	s := make([]string, 3) // 初始化一个长度为3的字符串数组
	fmt.Println("emp:", s)
	fmt.Println("len:", len(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

	s = append(s, "d")      // 追加一个元素
	s = append(s, "e", "f") // 追加两个元素
	fmt.Println("apd:", s)

	c := make([]string, len(s)) // 按原数组长度初始化一个新的字符串数组
	copy(c, s)                  // 从原数组中复制元素到新数组
	fmt.Println("cpy:", c)

	l := s[2:5] // 从index为2开始复制元素，复制长度为5
	fmt.Println("sl1:", l)

	l = s[:5] // 从index为0开始复制元素，复制长度为5
	fmt.Println("sl2:", l)

	l = s[2:] // 从index为2开始复制元素，直到最后一个
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} // 初始化一个动态长度数组
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) // 初始化一个二位数组
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
