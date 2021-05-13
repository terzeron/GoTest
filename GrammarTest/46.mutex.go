package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64 = 0
	var writeOps uint64 = 0

	for r := 0; r < 100; r++ {
		// 100개의 고루틴
		go func() {
			total := 0
			for {
				// map의 임의의 요소의 값을 읽어서 total에 누적합을 계산함
				key:=rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()

				// 읽기연산 횟수 카운터
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w:=0;w<10;w++ {
		// 10개의 고루틴
		go func() {
			for {
				// map의 임의의 요소에 임의의 값을 쓰기
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()

				// 쓰기연산 횟수 카운터
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

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
