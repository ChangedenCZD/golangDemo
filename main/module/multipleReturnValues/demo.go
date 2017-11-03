package multipleReturnValues
/**多值返回*/
import "fmt"

func vals() (int, int) {
	return 4, 14
}

func Run() {
	a, b := vals()
	fmt.Println(a, b)

	_, c := vals()
	fmt.Println(c)
}
