package main

import (
	"fmt"
	"time"
)

func main() {
	// 타이머는 채널을 만들어서 반환함
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C // blocking
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	// 타이머는 단순 sleep과 달리 중간에 취소가 가능함
	// 타이머 채널에 대한 blocking은 고루틴에서 비동기적으로 실행 중이기 때문에 메인 루틴은 다음의 취소 로직을 실행함
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
