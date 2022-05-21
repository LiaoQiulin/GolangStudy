package main

import (
	"flag"
	"fmt"
	"os"
)

// 一些命令行工具，如 go 工具或 git 有许多子命令，每个子命令都有自己的一组标志。
// 例如，go build 和 go get 是 go 工具的两个不同的子命令。 flag 包让我们可以轻松定义具有自己标志的简单子命令。
func main() {

	// 我们使用 NewFlagSet 函数声明一个子命令，然后继续为这个子命令定义新的标志。
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "enable")
	fooName := fooCmd.String("name", "", "name")

	// 对于不同的子命令，我们可以定义不同的支持标志。
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "level")

	// 该子命令应作为程序的第一个参数。
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// 检查调用了哪个子命令。
	switch os.Args[1] {

	// 对于每个子命令，我们解析它自己的标志并可以访问尾随位置参数。
	case "foo":
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		fmt.Println("  tail:", fooCmd.Args())
	case "bar":
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())
	default:
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	/*
		$ go build -o command-line-subcommands.exe

		先调用 foo 子命令。
		$ ./command-line-subcommands foo -enable -name=joe a1 a2
		subcommand 'foo'
			enable: true
			name: joe
			tail: [a1 a2]

		现在 bar。
		$ ./command-line-subcommands bar -level 8 a1
		subcommand 'bar'
			level: 8
			tail: [a1]

		但是 bar 不会接受 foo 的标志。
		$ ./command-line-subcommands bar -enable a1
		flag provided but not defined: -enable
		Usage of bar:
			-level int
				level
	*/
}
