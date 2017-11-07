package exitDemo

import (
	"fmt"
	"os"
)

/**退出引擎*/

func Run() {
	defer fmt.Println("!")

	os.Exit(3) // 状态码为3
}
