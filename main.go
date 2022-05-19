package main

import (
	"fmt"
	"sync"
	"time"
)

// 要等待多个 goroutine 完成，我们可以使用等待组。

// 这是我们将在每个 goroutine 中运行的函数。
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// 睡眠以模拟耗时任务。
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// 这个 WaitGroup 用于等待这里启动的所有 goroutine 完成。
	// 注意：如果一个 WaitGroup 显式传递给函数，它应该通过指针来完成。
	var wg sync.WaitGroup

	// 启动几个 goroutine 并为每个增加 WaitGroup 计数器。
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		// 避免在每个 goroutine 闭包中重复使用相同的 i 值。
		i := i

		// 将 worker 调用封装在一个闭包中，确保告诉 WaitGroup 该 worker 已完成。
		// 这样，worker 本身就不必知道其执行过程中涉及的并发原语。
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

	// 每次调用的工作人员启动和完成的顺序可能不同。
}
