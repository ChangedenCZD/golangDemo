package tickerDemo

import (
	"time"
	"fmt"
)

/**轮询，固定时间重复执行，interval*/

func Run() {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	timer := time.NewTimer(time.Millisecond * 1600)
	go func() {
		<-timer.C
		ticker.Stop()
		fmt.Println("Ticker stopped")
	}()
	fmt.Scanln()
}
