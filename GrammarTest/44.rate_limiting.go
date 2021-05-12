package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)
	// 채널에 5개의 데이터가 쌓여 있음

	// 티커 형태로 rate limiter를 생성함
	limiter := time.Tick(time.Millisecond * 500)

	// 채널에서 데이터를 하나씩 읽어오지만
	// limiter 때문에 ticker에서 지정한 주기대로 blocking되어 동작함
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 위의 rate limiter와 다르게 초반의 bursty한 요청을 처리하도록 개선함
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	// burstyLimiter에 3개의 timestamp 값이 쌓여 있음
	// 고루틴을 이용하여 비동기적으로 0.5초마다 timesmptamp 값을 쌓음
	go func() {
		for t := range time.Tick(time.Millisecond * 500) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// 처음 3개는 거의 시작하자마자 동시에 처리하지만 이후 데이터는 0.5초마다 하나씩 처리함
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
