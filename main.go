package main

import (
	"fmt"
	"time"
)

// Go 的 select 让你等待多个通道操作。将 goroutine 和通道与 select 相结合是 Go 的一个强大功能。
func main() {

	// 对于我们的示例，我们将跨两个通道进行选择。
	c1 := make(chan string)
	c2 := make(chan string)

	// 每个通道都会在一段时间后收到一个值，以模拟例如阻塞在并发 goroutine 中执行的 RPC 操作。
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 我们将使用 select 同时等待这两个值，并在每个值到达时打印它们。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
