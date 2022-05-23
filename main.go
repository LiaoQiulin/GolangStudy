package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读取和写入文件是许多 Go 程序所需的基本任务。首先，我们将看一些读取文件的示例。

// 读取文件需要检查大多数调用是否有错误。这个助手将简化我们下面的错误检查。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 也许最基本的文件读取任务是将文件的全部内容吞入内存。
	dat, err := os.ReadFile("./data.txt")
	check(err)
	fmt.Print(string(dat))

	// 您通常希望更好地控制文件的读取方式和读取部分。对于这些任务，首先打开文件以获取 os.File 值。
	f, err := os.Open("./data.txt")
	check(err)

	// 从文件的开头读取一些字节。最多允许阅读 5 个，但也要注意实际阅读了多少。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// 您还可以查找文件中的已知位置并从那里读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// io 包提供了一些可能对文件读取有帮助的函数。例如，使用 ReadAtLeast 可以更稳健地实现上述读取。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内置的倒带，但 Seek(0, 0) 完成了这一点。
	_, err = f.Seek(0, 0)
	check(err)

	// bufio 包实现了一个缓冲读取器，这对于它的许多小读取的效率以及它提供的其他读取方法可能很有用。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 完成后关闭文件（通常这会在使用 defer 打开后立即安排）。
	f.Close()
}
