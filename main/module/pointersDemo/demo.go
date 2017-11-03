package pointersDemo

import "fmt"

/**指针*/
func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}
func Run() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // 转换成指针
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
