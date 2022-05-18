package main

import "fmt"

func main() {

	// 通道是连接并发 goroutine 的管道。您可以将值从一个 goroutine 发送到通道，然后将这些值接收到另一个 goroutine。

	// 使用 make(chan val-type) 创建一个新频道。通道按它们传达的值进行分类。
	messages := make(chan string)

	// 使用通道 <- 语法将值发送到通道。在这里，我们从一个新的 goroutine 向我们上面创建的消息通道发送“ping”。
	go func() { messages <- "ping" }()

	// <-channel 语法从通道接收一个值。在这里，我们将收到我们上面发送的“ping”消息并将其打印出来。
	msg := <-messages
	fmt.Println(msg)

	// 默认情况下，发送和接收会阻塞，直到发送方和接收方都准备好。该属性允许我们在程序结束时等待“ping”消息，而无需使用任何其他同步
}
