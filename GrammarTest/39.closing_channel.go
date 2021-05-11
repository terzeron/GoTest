package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	//           ------jobs-------
	//             1 2 3
	// worker <- ----------------- <- main

	//         ------done------
	//            true
	// main <- ---------------- <- worker

	// worker
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 수신
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// 채널 종료
	// 채널을 닫지 않으면 계속 수신하겠다고 열어두는 것이므로
	// 송신쪽과 짝이 맞지 않아서 에러 발생함
	//close(jobs)
	fmt.Println("sent all jobs")

	// worker 대기 (동기화)
	<-done
}
