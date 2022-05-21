package main

import (
	"os"
	"os/exec"
	"syscall"
)

// 在前面的示例中，我们查看了生成外部进程。当我们需要一个正在运行的 Go 进程可以访问的外部进程时，我们会这样做。
// 有时我们只想用另一个（可能是非 Go）进程完全替换当前的 Go 进程。为此，我们将使用 Go 对经典 exec 函数的实现。
func main() {

	// 对于我们的示例，我们将执行 ls。
	// Go 需要一个指向我们要执行的二进制文件的绝对路径，因此我们将使用 exec.LookPath 来查找它（可能是 /bin/ls）。
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec 需要切片形式的参数（而不是一个大字符串）。
	// 我们将给出 ls 一些常见的参数。请注意，第一个参数应该是程序名称。
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec 还需要一组环境变量才能使用。这里我们只提供我们当前的环境。
	env := os.Environ()

	// 这是实际的 syscall.Exec 调用。
	// 如果这个调用成功，我们进程的执行将到此结束，并被 /bin/ls -a -l -h 进程代替。
	// 如果有错误，我们将得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

	// 请注意，Go 不提供经典的 Unix fork 函数。
	// 不过通常这不是问题，因为启动 goroutine、生成进程和执行进程涵盖了 fork 的大多数用例。
}
