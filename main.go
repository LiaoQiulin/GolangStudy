package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Go 有几个有用的函数来处理文件系统中的目录。

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 在当前工作目录中创建一个新的子目录。
	err := os.Mkdir("subdir", 0755)
	check(err)

	// 创建临时目录时，最好推迟删除它们。 os.RemoveAll 将删除整个目录树（类似于 rm -rf）。
	defer os.RemoveAll("subdir")

	// 帮助函数创建一个新的空文件。
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	createEmptyFile("subdir/file1")

	// 我们可以使用 MkdirAll 创建目录层次结构，包括父目录。这类似于命令行 mkdir -p。
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir 列出目录内容，返回一片 os.DirEntry 对象。
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Chdir 让我们改变当前的工作目录，类似于 cd。
	err = os.Chdir("subdir/parent/child")
	check(err)

	c, err = os.ReadDir(".")
	check(err)

	// 现在我们将在列出当前目录时看到 subdir/parent/child 的内容。
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// cd 回到我们开始的地方。
	err = os.Chdir("../../..")
	check(err)

	// 我们还可以递归地访问一个目录，包括它的所有子目录。 Walk 接受一个回调函数来处理访问的每个文件或目录。
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

// 对于 filepath.Walk 递归找到的每个文件或目录，都会调用 visit。
func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", p, info.IsDir())
	return nil
}
