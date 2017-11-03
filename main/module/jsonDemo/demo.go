package jsonDemo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Resp1 struct {
	Page   int
	Fruits []string
}
type Resp2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func Run() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"a", "b", "c"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"a": 1, "b": 2}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	/**以类中的成员名作为key*/
	resp1 := &Resp1{
		Page:   1,
		Fruits: []string{"a", "b", "c"}}
	resp1B, _ := json.Marshal(resp1)
	fmt.Println(string(resp1B))
	/**以类中的成员json别名作为key*/
	resp2 := &Resp2{
		1, []string{"d", "e", "f"}}
	resp2B, _ := json.Marshal(resp2)
	fmt.Println(string(resp2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)
	num := dat["num"].(float64)
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	fmt.Println(strs)

	str := `{"page":1,"fruits":["a","b"]}`
	res := Resp2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"a": 5, "b": 7}
	enc.Encode(d)
}
