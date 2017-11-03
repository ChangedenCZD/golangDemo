package mapsDemo

/**map*/
import (
	"fmt"
)

func Run() {
	m := make(map[string]int) // 初始化一个key为字符串，value为int的map

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2") // 删除k2的值
	fmt.Println("map:", m)

	prs1 := m["k2"]
	fmt.Println("prs1:", prs1)
	_, prs2 := m["k2"] // 忽略0或""
	fmt.Println("prs2:", prs2)

	n := map[string]int{"foo": 1, "bar": 2} // 初始化一个map，并设置相关key:value
	fmt.Println("map:", n)
}
