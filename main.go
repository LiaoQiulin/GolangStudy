package main

import (
	"fmt"
	"sync"
)

// 在前面的示例中，我们看到了如何使用原子操作来管理简单的计数器状态。
// 对于更复杂的状态，我们可以使用互斥锁跨多个 goroutine 安全地访问数据。

// Container 里有一个map字段；由于我们想从多个 goroutine 中同时更新它，我们添加了一个 Mutex 来同步访问。
// 请注意，互斥锁不能被复制，所以如果这个结构被传递，它应该通过指针来完成。
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// 在访问计数器之前锁定互斥锁；使用 defer 语句在函数末尾解锁它。
func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	// 请注意，互斥锁的零值可以按原样使用，因此这里不需要初始化。
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	// 此函数在循环中递增命名计数器。
	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	// 同时运行多个 goroutine；请注意，它们都访问同一个 Container，其中两个访问同一个计数器。
	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("a", 10000)
	go doIncrement("b", 10000)

	// 等待 goroutine 完成
	wg.Wait()
	fmt.Println(c.counters)
}
