package main

import (
	"fmt"
	"time"
)

// ticker适用于您想定期重复做某事。这是一个周期性滴答的代码示例，直到我们停止它
func main() {

	// Tickers 使用与计时器类似的机制：发送值的通道。
	// 在这里，我们将使用通道上的 select 内置函数来等待每 500 毫秒到达的值。
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// ticker 可以像计时器一样停止。一旦代码停止，它将不会在其通道上接收任何值。我们将在 1600 毫秒后停止。
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
