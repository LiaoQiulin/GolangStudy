package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 在前面的示例中，我们使用互斥锁的显式锁定来同步对多个 goroutine 的共享状态的访问。另一种选择是使用 goroutine 和通道的内置同步功能来实现相同的结果。
// 这种基于通道的方法与 Go 共享内存的想法一致，即通过通信并使每条数据恰好由 1 个 goroutine 拥有。

// 在这个例子中，我们的 state 将由一个 goroutine 拥有。
// 这将保证数据不会因并发访问而损坏。为了读取或写入该状态，其他 goroutine 将向拥有的 goroutine 发送消息并接收相应的回复。
// 这些 readOp 和 writeOp 结构体封装了这些请求以及拥有的 goroutine 响应的方式。
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 和之前一样，我们将计算我们执行了多少操作。
	var readOps uint64
	var writeOps uint64

	// 其他 goroutine 将使用读取和写入通道分别发出读取和写入请求。
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// 这是拥有状态的 goroutine，它是一个与前面示例一样的映射，但现在是有状态 goroutine 私有的。
	// 这个 goroutine 反复选择读取和写入通道，在请求到达时对其进行响应。
	// 通过首先执行请求的操作，然后在响应通道 resp 上发送一个值来指示成功（以及读取情况下的所需值）来执行响应。
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 这将启动 100 个 goroutine 以通过读取通道向拥有状态的 goroutine 发出读取。
	// 每次读取都需要构造一个 readOp，通过读取通道发送它，并通过提供的响应通道接收结果。
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 我们也使用类似的方法开始 10 次写入。
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	// 最后，捕获并报告操作计数。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
