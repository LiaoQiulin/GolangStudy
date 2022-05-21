package main

import (
	"fmt"
	"os"
)

// 使用 os.Exit 以给定状态立即退出。
func main() {

	// 使用 os.Exit 时不会运行 defers，因此永远不会调用此 fmt.Println。
	defer fmt.Println("!")

	// 以状态 3 退出。
	os.Exit(3)

	// 请注意，不像例如C、Go 不使用 main 的整数返回值来指示退出状态。
	// 如果你想以非零状态退出，你应该使用 os.Exit。
}
