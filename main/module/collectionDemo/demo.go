package collectionDemo

import (
	"fmt"
	"strings"
)

/**集合*/
/**获取目标位置*/
func Index(vs []string, target string) int {
	for i, v := range vs {
		if v == target {
			return i
		}
	}
	return -1
}

/**是否存在目标*/
func Include(vs []string, target string) bool {
	return Index(vs, target) >= 0
}

/**是否存在匹配，需要用匿名函数拓展*/
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

/**是否全部匹配，需要用匿名函数拓展*/
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

/**过滤器*/
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

/**用于创造新片段*/
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

/**去重*/
func Uniq(vs []string) []string {
	vsu := make([]string, 0)
	for i, v := range vs {
		if Index(vs, v) == i {
			vsu = append(vsu, v)
		}
	}
	return vsu
}

/**去重方式追加元素*/
func Append(vs []string, element string, isUniq bool) []string {
	vsa := vs
	if !(isUniq && Include(vs, element)) {
		vsa = append(vs, element)
	}
	return vsa
}

func Run() {
	var stringArray = []string{"peach", "apple", "pear", "apple", "plum"}

	fmt.Println(Index(stringArray, "pear"))

	fmt.Println(Include(stringArray, "grape"))

	fmt.Println(Any(stringArray, func(v string) bool {
		return strings.HasPrefix(v, "p") // 是否以p为前缀
	}))

	fmt.Println(All(stringArray, func(v string) bool {
		return strings.HasPrefix(v, "p") // 是否以p为前缀
	}))

	fmt.Println(Filter(stringArray, func(v string) bool {
		return strings.Contains(v, "e") //是否包含字符e
	}))

	fmt.Println(Map(stringArray, strings.ToUpper))

	fmt.Println(Uniq(stringArray))

	fmt.Println(Append(stringArray, "plum", false))

	fmt.Println(Append(stringArray, "peach", true))
}
