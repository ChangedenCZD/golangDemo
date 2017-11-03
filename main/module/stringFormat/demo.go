package stringFormat

import (
	"fmt"
	"os"
)

/**字符串格式化*/
type point struct {
	x, y int
}

func Run() {
	p := point{1, 2}
	/**只输入值*/
	fmt.Printf("%v\n", p)
	/**输出key:value*/
	fmt.Printf("%+v\n", p)
	/**包名.类名{key:value}*/
	fmt.Printf("%#v\n", p)
	/**包名.类名*/
	fmt.Printf("%T\n", p)
	/**输出bool*/
	fmt.Printf("%t\n", true)
	/**输出10进制整数*/
	fmt.Printf("%d\n", 123)
	/**输出二进制*/
	fmt.Printf("%b\n", 14)
	/**输出ascii对应编码的字符*/
	fmt.Printf("%c\n", 33)
	/**输出16进制*/
	fmt.Printf("%x\n", 456)
	/**输出浮点*/
	fmt.Printf("%f\n", 78.9)
	/**输出科学计算*/
	fmt.Printf("%e\n", 123400000.0)

	fmt.Printf("%E\n", 123400000.0)
	/**处理转义*/
	fmt.Printf("%s\n", "\"string\"")
	/**不处理转义*/
	fmt.Printf("%q\n", "\"string\"")
	/**输出hash码*/
	fmt.Printf("%x\n", "hex this")
	/**输出内存地址*/
	fmt.Printf("%p\n", &p)
	/**表格样式输出，cell size为6个字符*/
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	/**保留两位小数*/
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	/**居左显示*/
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")
	/**以string format 形式拼接字符串*/
	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)
	/**以标准错误输出内容*/
	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
