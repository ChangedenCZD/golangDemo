package timerDemo

import (
	"time"
	"fmt"
)

/**计时器*/
func Run() {
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println("before timer1.C", time.Now().UnixNano()/1000000)
	<-timer1.C
	fmt.Println("Timer 1 expired", time.Now().UnixNano()/1000000)

	timer2 := time.NewTimer(time.Second)
	fmt.Println("before timer2.C", time.Now().UnixNano()/1000000)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired", time.Now().UnixNano()/1000000)
	}()
	fmt.Println("before timer2.Stop()", time.Now().UnixNano()/1000000)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped", time.Now().UnixNano()/1000000)
	}
}
