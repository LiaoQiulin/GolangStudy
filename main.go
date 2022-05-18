package main

import (
	"fmt"
	"time"
)

// 我们可以使用通道来同步跨 goroutine 的执行。这是一个使用阻塞接收来等待 goroutine 完成的示例。
// 当等待多个 goroutine 完成时，您可能更喜欢使用 WaitGroup。

// 这是我们将在 goroutine 中运行的函数。 done 通道将用于通知另一个 goroutine 该函数的工作已完成。
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送一个值来通知我们已经完成。
	done <- true
}

func main() {

	// 启动一个worker goroutine，为其提供通知通道。
	done := make(chan bool, 1)
	go worker(done)

	// 阻塞，直到我们在频道上收到 worker 的通知。
	<-done
}
