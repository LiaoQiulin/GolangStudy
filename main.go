package main

import (
	"fmt"
	"io"
	"os/exec"
)

// 有时我们的 Go 程序需要生成其他非 Go 进程。
func main() {

	// 我们将从一个简单的命令开始，它不带参数或输入，只是将一些内容打印到标准输出。 e
	// xec.Command 助手创建一个对象来表示这个外部进程。
	dateCmd := exec.Command("go", "version")

	// Output 方法运行命令，等待它完成并收集它的标准输出。如果没有错误， dateOut 将保存带有日期信息的字节。
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> go version")
	fmt.Println(string(dateOut))

	// 如果执行命令时出现问题（例如错误的路径），输出和其他命令方法将返回 *exec.Error，
	// 如果命令运行但以非零返回码退出，则返回 *exec.ExitError
	_, err = exec.Command("go").Output()
	if err != nil {
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode())
		default:
			panic(err)
		}
	}

	// 接下来，我们将看一个稍微复杂的案例，我们将数据通过管道传输到其标准输入上的外部进程，并从标准输出收集结果。
	grepCmd := exec.Command("grep", "hello")

	// 在这里，我们显式地抓取输入/输出管道，启动进程，向其写入一些输入，读取结果输出，最后等待进程退出。
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()

	// 在上面的例子中我们省略了错误检查，但是你可以使用通常的 if err != nil 模式来检查所有的错误。
	// 我们也只收集 StdoutPipe 结果，但您可以以完全相同的方式收集 StderrPipe。
	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	// 请注意，在生成命令时，我们需要提供一个明确描述的命令和参数数组，而不是能够只传入一个命令行字符串。
	// 如果你想用一个字符串生成一个完整的命令，你可以使用 bash 的 -c 选项：
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}
