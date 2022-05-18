package main

import "fmt"

// 当使用通道作为函数参数时，您可以指定通道是否仅用于发送或接收值。这种特殊性增加了程序的类型安全性。

// 此 ping 函数仅接受用于发送值的通道。尝试在此通道上接收将是编译时错误。
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong 函数接受一个通道用于接收（pings）和第二个通道用于发送（pongs）。
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
