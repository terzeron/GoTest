package main

import "fmt"

func main() {
	// channel (like pipe)
	messages := make(chan string)

	// <-를 이용하여 메시지를 채널에 송신
	go func() { messages <- "ping" }()

	// <-를 이용하여 채널에서 메시지를 수신
	msg := <-messages
	fmt.Println(msg)
}
