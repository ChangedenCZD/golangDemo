package rangeDemo
/**范围取值*/
import "fmt"

func Run() {
	fmt.Println("usage 1: for index := range object")
	fmt.Println("usage 2: for index,item := range array")
	fmt.Println("usage 3: for key,value := range map")
	fmt.Println("usage 4: for index,ascii := range string")

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { // index位置用空白补位
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		fmt.Println("index:", i, num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("keys", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
