package main

import "fmt"

// Go 可以通过使用 recover 内置函数从 panic 中恢复。 recover 可以阻止 panic 中止程序并让它继续执行。

// 一个有用的例子：如果一个客户端连接出现严重错误，服务器不希望崩溃。
// 相反，服务器会想要关闭该连接并继续为其他客户端提供服务。事实上，这就是 Go 的 net/http 默认为 HTTP 服务器所做的。

// 这个函数 panic 。
func mayPanic() {
	panic("a problem")
}

// 必须在 defer 函数中调用 recover 。当封闭函数发生 panic 时，defer 将激活并且其中的恢复调用将捕获 panic。
func main() {

	defer func() {
		// recover 的返回值是调用 panic 时引发的错误。
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// 这段代码不会运行，因为 mayPanic 出现恐慌。 main 的执行在 panic 点停止，并在 defer 关闭时恢复。
	fmt.Println("After mayPanic()")
}
