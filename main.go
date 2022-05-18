package main

import "fmt"

func main() {

	// 默认情况下，通道是无缓冲的，这意味着如果有相应的接收 (<- chan) 准备好接收发送的值，它们将只接受发送 (chan <-)。缓冲通道接受有限数量的值，而这些值没有相应的接收器。

	// 在这里，我们创建了一个字符串通道，最多可缓冲 2 个值。
	messages := make(chan string, 2)

	// 因为这个通道是缓冲的，我们可以将这些值发送到通道中，而无需相应的并发接收。
	messages <- "buffered"
	messages <- "channel"

	// 稍后我们可以像往常一样接收这两个值。
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
