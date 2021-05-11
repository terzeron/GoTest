package main

import "fmt"

func main() {
	// 기본적으로 채널은 unbuffered 상태
	// 수신준비가 안 되어 있으면 block되거나 유실됨
	// 버퍼 크기 지정 가능
	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
