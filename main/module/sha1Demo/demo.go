package sha1Demo

import (
	"crypto/sha1"
	"fmt"
)

/**生成sha1和hash码*/
func Run() {
	str := "sha1 this string"
	/**创建hash对象*/
	hashStr := sha1.New()
	/**将字符串转换为字节数组写入hash对象*/
	hashStr.Write([]byte(str))
	byteSlice := hashStr.Sum(nil)
	fmt.Println(str)
	fmt.Printf("%x\n", byteSlice)
}
