package panicDemo

import (
	"os"
)

/**抛出意外的错误*/

func Run() {
	//panic("a problem")

	_, err := os.Create("/a/b")
	if err != nil {
		panic(err)
	}
}
