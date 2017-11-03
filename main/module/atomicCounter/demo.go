package atomicCounter

import (
	"sync/atomic"
	"time"
	"fmt"
)

/**原子计数器*/

func Run() {
	var ops uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
