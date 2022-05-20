package main

import (
	"bufio"
	"fmt"
	"os"
)

// 在 Go 中写文件遵循与我们之前看到的读文件模式类似的模式。
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 首先，这是将字符串（或只是字节）转储到文件中的方法。
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("../dat1", d1, 0644)
	check(err)

	// 要进行更精细的写入，请打开一个文件进行写入。
	f, err := os.Create("../dat2")
	check(err)

	// 打开文件后立即推迟关闭是惯用的。
	defer f.Close()

	// 您可以按预期写入字节切片。
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// 也可以使用 WriteString。
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// 调用 Sync 将写入刷新到稳定存储。
	f.Sync()

	// 除了我们之前看到的缓冲读取器之外，bufio 还提供缓冲写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用 Flush 确保所有缓冲操作都已应用于底层写入器。
	w.Flush()

}
