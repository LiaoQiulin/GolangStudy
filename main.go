package main

import (
	"fmt"
	"os"
	"strings"
)

// 环境变量是一种将配置信息传递给 Unix 程序的通用机制。让我们看看如何设置、获取和列出环境变量。
func main() {

	// 要设置键/值对，请使用 os.Setenv。要获取键的值，请使用 os.Getenv。
	// 如果密钥不存在于环境中，这将返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// 使用 os.Environ 列出环境中的所有键/值对。这将返回 KEY=value 形式的字符串切片。
	// 您可以使用 strings.SplitN 它们来获取键和值。在这里，我们打印所有的键。
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

}
