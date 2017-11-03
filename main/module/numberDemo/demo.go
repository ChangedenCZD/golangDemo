package numberDemo

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

/**数字类操作*/

func Run() {
	/**获取0~99之间的随机数*/
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	/**获取0~1之间的随机数*/
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	seed1 := rand.NewSource(time.Now().UnixNano())
	random1 := rand.New(seed1)
	fmt.Print(random1.Intn(100), ",")
	fmt.Print(random1.Intn(100))
	fmt.Println()

	/**设置随机种子*/
	seed2 := rand.NewSource(42)
	random2 := rand.New(seed2)
	fmt.Print(random2.Intn(100), ",")
	fmt.Print(random2.Intn(100))
	fmt.Println()
	/**种子与 seed2 一致，因此结果是一致的*/
	seed3 := rand.NewSource(42)
	random3 := rand.New(seed3)
	fmt.Print(random3.Intn(100), ",")
	fmt.Print(random3.Intn(100))
	fmt.Println()

	fmt.Println("----------------")
	/**数字与字符串转换*/
	/**转换成64位浮点*/
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	/**转换成long类型，默认值为0*/
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)
	/**将16进制转为10进制*/
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	/**转换为无符long类型*/
	u, _ := strconv.ParseUint("-789", 0, 64)
	fmt.Println(u)

	/**转为integer类型*/
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	/**字符串不符合数字格式*/
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
