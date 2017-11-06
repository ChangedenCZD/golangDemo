package spawnProcesses

import (
	"os/exec"
	"fmt"
	"os"
	"syscall"
	"io/ioutil"
)

/**生产执行进程*/
var p = fmt.Println

func Run() {
	dateCmd := exec.Command("npm")
	dateOut, err := dateCmd.Output()
	checkError(err)
	p("> npm")
	p(string(dateOut))

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	p("> grep hello")
	p(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	checkError(err)
	p("> ls -a -l -h")
	p(string(lsOut))

	binary, lookErr := exec.LookPath("npm")
	checkError(lookErr)
	args := []string{"npm", "-version"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	checkError(execErr)
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
