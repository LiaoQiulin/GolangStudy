package main

import "fmt"

// 我们可以使用 range 迭代从通道接收到的值。
func main() {

	// 我们将在队列通道中迭代 2 个值。
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// range 会迭代从 queue 中接收到每个元素。因为我们关闭了上面的通道，所以迭代在接收到 2 个元素后终止。
	for elem := range queue {
		fmt.Println(elem)
	}
}
