package sortDemo

import (
	"sort"
	"fmt"
)

// 为字符串数组创建一个别名ByLength
type ByLength []string

/**
实现sort.Interface 中的Len方法
排序前用于判断对象长度
*/
func (s ByLength) Len() int {
	fmt.Println("sort.Interface.Len", len(s))
	return len(s)
}

/**
实现sort.Interface 中的Swap方法
用于转换位置
*/
func (s ByLength) Swap(l, r int) {
	fmt.Println("sort.Interface.Swap", l, r)
	s[l], s[r] = s[r], s[l]
}

/**
实现sort.Interface 中的Less方法
判断是否需要转换位置
*/
func (s ByLength) Less(l, r int) bool {
	fmt.Println("sort.Interface.Less", l, r)
	return len(s[l]) < len(s[r])
}

/**排序*/
func Run() {
	stringArray := []string{"c", "a", "b"}
	sort.Strings(stringArray)
	fmt.Println("Strings:", stringArray)

	intArray := []int{44, 4, 14}
	sort.Ints(intArray)
	fmt.Println("Ints:", intArray)

	s := sort.IntsAreSorted(intArray) // 判断是否已排序
	fmt.Println("Sorted:", s)

	fruitArray := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruitArray)) // 按文本长度排序
	fmt.Println(fruitArray)
}
