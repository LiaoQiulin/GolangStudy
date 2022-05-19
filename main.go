package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	// 我们将使用一个无符号整数来表示我们的（始终为正的）计数器
	var ops uint64

	// WaitGroup 将帮助我们等待所有 goroutine 完成它们的工作。
	var wg sync.WaitGroup

	// 我们将启动 50 个 goroutine，每个 goroutine 将计数器递增 1000 次。
	for i := 0; i < 50; i++ {
		wg.Add(1)

		// 为了原子地增加计数器，我们使用 AddUint64，使用 & 语法给它我们的操作计数器的内存地址作为入参。
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {

				// ops++
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	// 等到所有的 goroutine 都完成。
	wg.Wait()

	// 现在访问 ops 是安全的，因为我们知道没有其他 goroutine 正在写入它。使用 atomic.LoadUint64 等函数也可以在更新原子时安全地读取它们。
	fmt.Println("ops:", ops)

	// 我们预计将完成 50,000 次操作。如果我们使用非原子操作++来增加计数器，我们可能会得到一个不同的数字，在运行之间改变，因为 goroutines 会相互干扰。
	// 此外，当使用 -race 标志运行时，我们会遇到数据竞争失败。
}
