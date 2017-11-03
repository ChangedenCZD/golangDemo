package arrayDemo
/**数组*/
import "fmt"

func Run() {
	var a [5]int // 初始化长度为5的int类型的数组
	fmt.Println("emp:", a)

	a[4] = 100 // 设置index为4的值
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a)) // 获取数据长度

	b := [6]int{1, 2, 3, 4, 5} // 初始化长度为6的int类型的数组，且初始化各个值
	fmt.Println("dcl:", b)

	var twoD [2][3]int // 初始化一个二位数组
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}
