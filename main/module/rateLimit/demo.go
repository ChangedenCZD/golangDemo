package rateLimit

import (
	"time"
	"fmt"
)

/**限速器机制*/
func current() int64 {
	return time.Now().UnixNano() / 1000000
}
func Run() {
	/**
	* 基础模式
	*/
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 200)
	startTime := current()
	for req := range requests {
		<-limiter
		currentTime := current()
		fmt.Println("request I", req, currentTime-startTime)
		startTime = currentTime
	}

	/**
	* 设立缓冲区
	*/

	burstyLimiter := make(chan time.Time, 3) // 缓冲区为3个scheme

	for i := 0; i < 3; i++ { // 设置前3个缓冲区为0延时
		burstyLimiter <- time.Now()
	}

	go func() { // 当burstyLimiter缓冲区不满时补入限速器
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	startTime = current()
	for req := range burstyRequests {
		<-burstyLimiter
		currentTime := current()
		fmt.Println("request II", req, currentTime-startTime)
		startTime = currentTime
	}
}
