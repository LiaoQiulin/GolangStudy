package main

import "os"

// panic 通常意味着出现意外错误。
// 大多数情况下，我们使用它来快速解决在正常操作期间不应该发生的错误，或者我们没有准备好优雅处理的错误。
func main() {

	// 我们将在整个站点中使用 panic 来检查意外错误。这是该网站上唯一旨在 panic 的程序。
	panic("a problem")

	// 如果一个函数返回一个我们不知道如何（或想要）处理的错误值，panic 的一个常见用途是中止。
	// 如果我们在创建新文件时遇到意外错误，这是一个 panic 的例子
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
