package main

import (
	"fmt"
	"os"
)

// 命令行参数是参数化程序执行的常用方法。
// 例如， go run hello.go 使用 go 程序的 run 和 hello.go 参数。

func main() {

	// os.Args 提供对原始命令行参数的访问。请注意，此切片中的第一个值是程序的路径，os.Args[1:] 保存程序的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 您可以获得具有正常索引的单个参数。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// 要试验命令行参数，最好先使用 go build 构建二进制文件。
	/*
		$ go build command-line-arguments.go
		$ ./command-line-arguments a b c d
		[./command-line-arguments a b c d]
		[a b c d]
		c
	*/
}
