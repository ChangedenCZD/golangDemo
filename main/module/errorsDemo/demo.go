package errorsDemo

import (
	"errors"
	"fmt"
)

/**错误，异常*/

func f1(arg int) (int, error) {
	if arg == 42 { // 如果arg等于42则报错
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type argError struct {
	//封装一个错误类
	arg  int
	prob string
}

func (e *argError) Error() string { // 输出错误信息
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 { // 如果arg等于42则报错
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
func Run() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
