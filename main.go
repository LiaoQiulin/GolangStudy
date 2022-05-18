package main

import (
	"fmt"
	"time"
)

// 超时对于连接到外部资源或需要限制执行时间的程序很重要。多亏了 channels 和 select ，在 Go 中实现超时变得简单而优雅。
func main() {

	// 对于我们的示例，假设我们正在执行一个外部调用，该调用在 2 秒后在通道 c1 上返回其结果。
	// 请注意，通道是缓冲的，因此 goroutine 中的发送是非阻塞的。
	// 这是防止 goroutine 泄漏的常见模式，以防从不读取通道。
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// 这是实现超时的 select。
	// res := <-c1 等待结果，<-time.After 等待超时 1s 后发送的值。
	// 由于 select 会处理第一个准备好的接收，因此如果操作花费的时间超过允许的 1，我们将采用超时 case 。
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// 如果我们允许更长的超时时间为 3 秒，那么来自 c2 的接收将成功，我们将打印结果。
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}
