package main

import "fmt"

// 通道上的基本发送和接收是阻塞的。
// 但是，我们可以使用带有 default 子句的 select 来实现非阻塞发送、接收，甚至是非阻塞多路选择。
func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 这是一个非阻塞接收。
	// 如果消息上有可用的值，则 select 将采用具有该值的 <-messages 大小写。
	// 如果不是，它将立即采用默认情况。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送的工作方式类似。
	// 这里 msg 不能发送到消息通道，因为通道没有缓冲区，也没有接收器。因此选择默认情况。
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 我们可以在 default 子句之上使用多种情况来实现多路非阻塞选择。
	// 在这里，我们尝试对消息和信号进行非阻塞接收。
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
