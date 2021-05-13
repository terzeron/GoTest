package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readMsgStruct struct {
	key  int
	resp chan int
}

type writeMsgStruct struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var readOps uint64 = 0
	var writeOps uint64 = 0

	readChannel := make(chan *readMsgStruct)
	writeChannel := make(chan *writeMsgStruct)

	go func() {
		// state는 하나의 고루틴이 독점 사용함 (lock 불필요함)
		var state = make(map[int]int)
		for {
			select {
			// 읽기채널로 메시지가 오면 응답채널로 map의 특정 키가 가리키는 요소의 값을 보냄
			case readMsg := <-readChannel:
				readMsg.resp <- state[readMsg.key]
			// 쓰기채널로 메시지가 오면 데이터를 map에 저장하고 응답채널로 true를 보냄
			case writeMsg := <-writeChannel:
				state[writeMsg.key] = writeMsg.val
				writeMsg.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		// 100개의 읽기 고루틴
		go func() {
			for {
				readMsg := &readMsgStruct{
					key:  rand.Intn(5),
					resp: make(chan int)}
				readChannel <- readMsg
				<-readMsg.resp

				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		// 10개의 쓰기 고루틴
		go func() {
			for {
				writeMSg := &writeMsgStruct{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writeChannel <- writeMSg
				<-writeMSg.resp

				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
