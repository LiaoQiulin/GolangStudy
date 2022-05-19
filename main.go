package main

import (
	"fmt"
	"os"
)

// Defer 用于确保在程序执行的后期执行函数调用，通常用于清理目的。
//  defer 经常用于例如 ensure  和 finally  在其他语言使用的地方。

// 假设我们想创建一个文件，写入它，然后在完成后关闭。下面是我们如何使用 defer 来做到这一点。
func main() {

	// 在使用 createFile 获取文件对象后，我们立即使用 closeFile 推迟关闭该文件。
	// 这将在 writeFile 完成后，在封闭函数（main）的末尾执行。
	f := createFile("D://defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")

}

// 关闭文件时检查错误很重要，即使在延迟函数中也是如此。
func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
