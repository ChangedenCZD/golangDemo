package regularExpression

import (
	"regexp"
	"fmt"
	"bytes"
)

/**正则表达式*/
func Run() {
	pattern := "p([a-z]+)ch"
	/**寻找前缀为p，后缀为ch，且中间至少有一个小写字母的字符串*/
	match, _ := regexp.MatchString(pattern, "peach")
	fmt.Println(match)

	r, _ := regexp.Compile(pattern)
	fmt.Println(r.MatchString("peach"))
	/**从左至右寻找匹配的字符串*/
	fmt.Println(r.FindString("peach punch"))
	/**从左至右寻找匹配的字符出现始末位置*/
	fmt.Println(r.FindStringIndex("peach punch"))
	/**寻找与副正则 [a-z]+ 也匹配的字符串*/
	fmt.Println(r.FindStringSubmatch("peach punch"))

	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	/**寻找所有与之匹配的字符串*/
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	/**寻找前两个匹配*/
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	fmt.Println(r.Match([]byte("peach")))

	r = regexp.MustCompile(pattern)
	fmt.Println(r)

	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	fmt.Println(string(r.ReplaceAllFunc([]byte("a peach"), bytes.ToUpper)))
}
