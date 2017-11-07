package signalDemo

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

/**系统信号*/
var p = fmt.Println

func Run() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		p()
		p(sig)
		done <- true
	}()

	p("awaiting signal")
	<-done
	p("exiting")
}
