package main

import "fmt"

// 关闭通道表示将不再在其上发送任何值。这对于向通道的接收者传达完成信息很有用。
func main() {

	// 在本例中，我们将使用 jobs 通道将要完成的工作从 main() goroutine 传递到 worker goroutine。
	// 当工人没有更多 jobs 时，我们将关闭工作通道。
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 这是 worker goroutine。它反复通过 j, more := <-jobs 从 jobs 接收 job。
	// 在这种特殊的 2-value 形式的接收中，如果作业已关闭并且通道中的所有值都已被接收，则 more 值将为 false。
	// 当我们完成所有工作时，我们使用它来通知完成。
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

	// 这会通过 jobs 通道向工作人员发送 3 个jobs，然后将其关闭。
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
