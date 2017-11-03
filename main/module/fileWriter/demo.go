package fileWriter

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)

/**文件写入*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var target1 = "F:/Changeden/web_changeden_net/golangDemo/temp1.txt"
var target2 = "F:/Changeden/web_changeden_net/golangDemo/temp2.txt"

func Run() {
	data1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile(target1, data1, 0644)
	check(err)

	file, err := os.Create(target2)
	check(err)
	defer file.Close()

	data2 := []byte{115, 111, 109, 101, 10}
	n2, err := file.Write(data2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := file.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	file.Sync()

	writer := bufio.NewWriter(file)
	n4, err := writer.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	writer.Flush()
}
