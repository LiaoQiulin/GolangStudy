package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 在整个程序执行过程中，我们经常希望创建程序退出后不需要的数据。
// 临时文件和目录对此很有用，因为它们不会随着时间的推移污染文件系统。

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 创建临时文件的最简单方法是调用 os.CreateTemp。它创建一个文件并打开它以进行读取和写入。
	// 我们提供 "" 作为第一个参数，因此 os.CreateTemp 将在我们的操作系统的默认位置创建文件。
	f, err := os.CreateTemp("", "sample.js")
	check(err)

	// 显示临时文件的名称。在基于 Unix 的操作系统上，该目录可能是 /tmp。
	// 文件名以作为 os.CreateTemp 的第二个参数给出的前缀开头，其余部分是自动选择的，以确保并发调用将始终创建不同的文件名。
	fmt.Println("Temp file name:", f.Name())

	// 完成后清理文件。操作系统可能会在一段时间后自行清理临时文件，但最好明确地执行此操作。
	defer os.Remove(f.Name())

	// 我们可以将一些数据写入文件。
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// 如果我们打算写很多临时文件，我们可能更喜欢创建一个临时目录。 os.MkdirTemp 的参数与 CreateTemp 的参数相同，但它返回的是目录名而不是打开的文件。
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	defer os.RemoveAll(dname)

	// 现在我们可以通过在临时目录前面加上临时文件名来合成临时文件名。
	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
