package main

import (
	"fmt"
	"time"
)

// 在本例中，我们将了解如何使用 goroutine 和通道实现工作池。

// 这是worker，我们将在其中运行几个并发实例。
// 这些工作人员将在 jobs 通道上接收工作，并在 results 上发送相应的结果。
// 我们将为每个作业睡一秒钟以模拟一项耗时的任务。
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// 为了使用我们的工人池，我们需要向他们发送工作并收集他们的结果。我们为此制作了 2 个通道。
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 这启动了 3 个工作人员，最初被阻止是因为还没有工作。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 在这里，我们发送 5 个工作，然后关闭该通道以表明这就是我们所有的工作。
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后我们收集所有工作的结果。这也确保了工作者 goroutines 已经完成。等待多个 goroutine 的另一种方法是使用 WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}

	// 我们的运行程序显示了由不同工人执行的 5 个作业。该程序只需要大约 2 秒，尽管总共做了大约 5 秒的工作，因为有 3 名工人同时操作。
}
