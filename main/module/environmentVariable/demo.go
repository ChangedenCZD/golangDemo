package environmentVariable

import (
	"os"
	"fmt"
	"strings"
)

var p = fmt.Println
/**环境变量*/
func Run() {
	os.Setenv("FOO", "1")
	p("FOO:", os.Getenv("FOO"))
	p("BAR:", os.Getenv("BAR"))

	p()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		p(pair)
	}
}
