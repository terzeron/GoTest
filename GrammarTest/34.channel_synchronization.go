package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 1초 대기후 true 메시지 전송
	done <- true
}

func main() {
	// 채널 생성
	done := make(chan bool, 1)
	// 고루틴 실행
	go worker(done)
	// 수신 대기
	<-done
}
