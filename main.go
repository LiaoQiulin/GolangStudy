package main

import (
	"fmt"
	"time"
)

// 速率限制是控制资源利用率和保持服务质量的重要机制。
// Go 优雅地支持 goroutines、channels 和 tickers 的速率限制。
func main() {

	// 首先我们来看看基本的速率限制。
	// 假设我们想限制对传入请求的处理。我们将通过同名通道处理这些请求。
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// // 此限制器通道将每 200 毫秒接收一个值。这是我们速率限制方案中的调节器。
	limiter := time.Tick(200 * time.Millisecond)

	// 通过在处理每个请求之前阻止来自限制器通道的接收，我们将自己限制为每 200 毫秒 1 个请求。
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 我们可能希望在我们的速率限制方案中允许短暂的请求突发，同时保留整体速率限制。
	// 我们可以通过缓冲我们的限制器通道来实现这一点。这个 burstyLimiter 通道将允许最多 3 个事件的突发。
	burstyLimiter := make(chan time.Time, 3)

	// 填充通道以表示允许的突发。
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 200 毫秒，我们将尝试向 burstyLimiter 添加一个新值，最大限制为 3。
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// 现在模拟另外 5 个传入请求。其中前 3 个将受益于 burstyLimiter 的突发能力。
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}

	// 对于第二批请求，由于突发速率限制，我们立即为前 3 个请求提供服务，然后为其余 2 个请求提供约 200 毫秒的延迟。
}
