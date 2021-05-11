package main

import "fmt"

// 채널이 수신용인지 송신용인지 타입을 정하는 방법

// pings: 송신 채널(chan<-)
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings: 수신 채널
// pongs: 송신 채널(<-chan)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
