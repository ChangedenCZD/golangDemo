package outerInput

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)

/**外部输入*/
func Run() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("请输入内容：")
	for scanner.Scan() {
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
		fmt.Println("请输入内容：")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
