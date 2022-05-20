package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

// filepath 包提供了解析和构造文件路径的功能，可以在操作系统之间移植；
// 例如，Linux 上的 dir/file 与 Windows 上的 dir\file。
func main() {

	// Join 应该用于以可移植的方式构造路径。它接受任意数量的参数并从它们构造一个分层路径。
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p)

	// 您应该始终使用 Join 而不是手动连接 /s 或 \s。除了提供可移植性之外，Join 还将通过删除多余的分隔符和目录更改来规范路径。
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	// Dir 和 Base 可用于拆分目录和文件的路径。或者，Split 将在同一个调用中返回两者。
	fmt.Println("Dir(p):", filepath.Dir(p))
	fmt.Println("Base(p):", filepath.Base(p))

	// 我们可以检查一个路径是否是绝对的。
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"

	// 一些文件名在点之后有扩展名。我们可以使用 Ext 将扩展名从这些名称中拆分出来。
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// 要查找已删除扩展名的文件名，请使用 strings.TrimSuffix。
	fmt.Println(strings.TrimSuffix(filename, ext))

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// Rel 找到基础和目标之间的相对路径。如果目标不能相对于基数，则返回错误。
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
}
