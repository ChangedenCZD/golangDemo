package stringsUtils

/**strings 工具类*/

import (
	s "strings"
	"fmt"
)

var p = fmt.Println

func Run() {
	p("Contains:", s.Contains("Test", "es"))
	/**strings.Count计算字符重复数，不忽略大小写*/
	p("Count:", s.Count("Test", "t"))
	/**strings.HasPrefix是否以某字符串为前缀，类似startsWith*/
	p("HasPrefix:", s.HasPrefix("Test", "Te"))
	/**strings.HasSuffix是否以某字符串为后缀，类似endsWith*/
	p("HasSuffix:", s.HasSuffix("Test", "st"))
	p("Index:", s.Index("Test", "e"))
	p("Join:", s.Join([]string{"a", "b"}, ":"))
	p("Repeat:", s.Repeat("14", 4))
	/**strings.Replace(原字符串,要替换的字符,新字符,替换数量：-1为全部)*/
	p("Replace1:", s.Replace("foo", "o", "0", -1))
	p("Replace2", s.Replace("foo", "o", "0", 1))
	p("Split:", s.Split("abcde", ""))
	p("ToLower:", s.ToLower("DEMO"))
	p("ToUpper:", s.ToUpper("demo"))
	p("Len:", len("strings"))
	p("Char:", "golang"[4])
}
