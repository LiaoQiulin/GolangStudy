package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 行过滤器是一种常见的程序类型，它读取标准输入上的输入，对其进行处理，然后将一些派生结果打印到标准输出。
// grep 和 sed 是常用的行过滤器。

// 这是 Go 中的一个示例行过滤器，它写入所有输入文本的大写版本。
// 您可以使用此模式编写自己的 Go 行过滤器。

func main() {

	// 用缓冲扫描器包装无缓冲的 os.Stdin 为我们提供了一种方便的 Scan 方法，可以将扫描器推进到下一个令牌；这是默认扫描仪的下一行。
	scanner := bufio.NewScanner(os.Stdin)

	// 文本从输入返回当前标记，这里是下一行。
	for scanner.Scan() {

		ucl := strings.ToUpper(scanner.Text())

		// 写出大写的行。
		fmt.Println(ucl)
	}

	// 在扫描期间检查错误。文件结尾是预期的，Scan 不会将其报告为错误。
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
