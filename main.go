package main

import (
	"flag"
	"fmt"
)

// 命令行标志是为命令行程序指定选项的常用方法。例如，在 wc -l 中，-l 是一个命令行标志。

func main() {

	// 基本标志声明可用于字符串、整数和布尔选项。这里我们声明一个带有默认值“foo”和简短描述的字符串标志字。
	// 这个 flag.String 函数返回一个字符串指针（不是字符串值）；我们将在下面看到如何使用这个指针。
	wordPtr := flag.String("word", "foo", "a string")

	// 这声明了 numb 和 fork 标志，使用与word 标志类似的方法。
	numbPtr := flag.Int("numb", 42, "an int")
	forkPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 声明所有标志后，调用 flag.Parse() 执行命令行解析。
	flag.Parse()

	// 在这里，我们将只转储已解析的选项和任何尾随位置参数。请注意，我们需要取消引用指针，例如*wordPtr 获取实际的选项值。
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	/*
		要试验命令行标志程序，最好先编译它，然后直接运行生成的二进制文件。
		$ go build command-line-flags.go

		首先给它所有标志的值来尝试构建的程序。
		$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
		word: opt
		numb: 7
		fork: true
		svar: flag
		tail: []

		请注意，如果您省略标志，它们会自动采用默认值。
		$ ./command-line-flags -word=opt
		word: opt
		numb: 42
		fork: false
		svar: bar
		tail: []

		可以在任何标志之后提供尾随位置参数。
		$ ./command-line-flags -word=opt a1 a2 a3
		word: opt
		...
		tail: [a1 a2 a3]

		请注意，标志包要求所有标志出现在位置参数之前（否则标志将被解释为位置参数）。
		$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
		word: opt
		numb: 42
		fork: false
		svar: bar
		tail: [a1 a2 a3 -numb=7]

		使用 -h 或 --help 标志为命令行程序自动生成帮助文本。
		$ ./command-line-flags -h
		Usage of ./command-line-flags:
		-fork=false: a bool
		-numb=42: an int
		-svar="bar": a string var
		-word="foo": a string

		如果您提供了一个未指定给 flag 包的标志，程序将打印一条错误消息并再次显示帮助文本。
		$ ./command-line-flags -wat
		flag provided but not defined: -wat
		Usage of ./command-line-flags:
		...
	*/

}
