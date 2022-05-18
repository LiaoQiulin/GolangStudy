package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Microsecond)
	}
}

func main() {

	// 假设我们有一个函数调用 f(s)。这是我们如何以通常的方式调用它，同步运行它。
	f("direct")

	// 要在 goroutine 中调用此函数，请使用 go f(s)。这个新的 goroutine 将与调用的 goroutine 同时执行。
	go f("goroutine")

	// 您还可以为匿名函数调用启动一个 goroutine。
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Println("sleep")
	// 我们的两个函数调用现在在不同的 goroutine 中异步运行。等待他们完成
	time.Sleep(time.Second)
	fmt.Println("done")

	// 我们运行这个程序时，我们首先看到阻塞调用的输出，然后是两个 goroutine 的输出。
	// goroutines 的输出可能是交错的，因为 goroutines 是由 Go 运行时并发运行的。
}
