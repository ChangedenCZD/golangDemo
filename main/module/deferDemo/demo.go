package deferDemo

import (
	"os"
	"bufio"
)

/**延时处理，最后执行，相当于finally*/

func Run() {
	path := "F:/Changeden/web_changeden_net/golangDemo/temp.txt"
	file := openFile(path)
	if file == nil {
		file = createFile(path)
	}
	defer closeFile(file)
	writeFile(file, "new data")
}

/**打开文件
追加内容，文件不存在则创建
os.O_RDONLY：只读
os.O_WRONLY：只写
os.O_CREATE：创建：如果指定文件不存在，就创建该文件。
os.O_TRUNC：截断：如果指定文件已存在，就将该文件的长度截为0
设置操作权限为0666,
*/
func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return nil
	}
	return f
}

/**创建文件*/
func createFile(path string) *os.File {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

/**写入文件*/
func writeFile(file *os.File, data string) {
	//readBuffer := bufio.NewReader(file)
	//oldData, _ := readBuffer.ReadString('\n')
	buffer := bufio.NewWriter(file) // 写入文件
	buffer.WriteString("\n" + data)
	buffer.Flush()
}

func closeFile(file *os.File) {
	file.Close()
}
