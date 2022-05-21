package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 有时我们希望我们的 Go 程序能够智能地处理 Unix 信号。
// 例如，我们可能希望服务器在收到 SIGTERM 时正常关闭，或者希望命令行工具在收到 SIGINT 时停止处理输入。
// 这是在 Go 中使用通道处理信号的方法。
func main() {

	// Go 信号通知通过在通道上发送 os.Signal 值来工作。
	// 我们将创建一个频道来接收这些通知。请注意，此通道应被缓冲。
	sigs := make(chan os.Signal, 1)

	// signal.Notify 注册给定通道以接收指定信号的通知。
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 我们可以在 main 函数中接收来自 sigs 的信息，但让我们看看如何在单独的 goroutine 中完成，以演示更真实的优雅关闭场景。
	done := make(chan bool, 1)

	// 这个 goroutine 对信号执行阻塞接收。当它得到一个时，它会打印出来，然后通知程序它可以完成。
	go func() {

		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 程序将在这里等待，直到它获得预期的信号（如上面的 goroutine 在完成时发送一个值所示），然后退出。
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	// 当我们运行这个程序时，它会阻塞等待信号。
	// 通过键入 ctrl-C（终端显示为 ^C），我们可以发送一个 SIGINT 信号，导致程序打印中断然后退出。
}
