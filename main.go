package main

import (
	"fmt"
	"time"
)

// 我们经常希望在未来的某个时间点执行 Go 代码，或者在某个时间间隔重复执行。
// Go 的内置 timer  和 ticker 功能使这两项任务变得简单。
// 我们将首先学习 timer  ，后面再学习 ticker。
func main() {

	// 计时器代表未来的单个事件。您告诉计时器您要等待多长时间，它提供了一个届时将通知的通道。此计时器将等待 2 秒。
	timer1 := time.NewTimer(2 * time.Second)

	// <-timer1.C 在定时器的通道 C 上阻塞，直到它发送一个指示定时器触发的值。
	<-timer1.C
	fmt.Println("Timer 1 fired")

	// 如果您只是想等待，您可以使用 time.Sleep。
	// 计时器可能有用的一个原因是您可以在计时器触发之前取消它。这是一个例子。
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// 给 timer2 足够的时间来触发它，如果它会触发的话，以表明它实际上已经停止了
	time.Sleep(3 * time.Second)
}
