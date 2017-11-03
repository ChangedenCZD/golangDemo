package fileReader

import (
	"io/ioutil"
	"fmt"
	"os"
	"io"
	"bufio"
)

/**文件读取*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var target = "F:/Changeden/web_changeden_net/golangDemo/temp.txt"

func Run() {
	/**获取文件中所有内容*/
	dat, err := ioutil.ReadFile(target)
	check(err)
	fmt.Print(string(dat))
	fmt.Println("\n------")

	file, err := os.Open(target)
	check(err)
	/**获取文件前5个内容*/
	buffer1 := make([]byte, 5)
	n1, err := file.Read(buffer1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(buffer1))
	fmt.Println("------")

	/**从第6个开始获取2个内容*/
	o2, err := file.Seek(6, 0)
	check(err)
	buffer2 := make([]byte, 2)
	n2, err := file.Read(buffer2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(buffer2))
	fmt.Println("------")

	o3, err := file.Seek(6, 0)
	check(err)
	buffer3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(file, buffer3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(buffer3))
	fmt.Println("------")

	/**获取前5个内容*/
	_, err = file.Seek(0, 0)
	check(err)
	reader := bufio.NewReader(file)
	buffer4, err := reader.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(buffer4))

	file.Close()
}
