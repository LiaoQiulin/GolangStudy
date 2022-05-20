package main

import (
	"fmt"
	"strconv"
)

// 从字符串中解析数字是许多程序中基本但常见的任务。这是在 Go 中的操作方法。
func main() {

	// 使用 ParseFloat，这 64 表示要解析多少位精度。
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 对于 ParseInt，0 表示从字符串推断基数。 64 要求结果适合 64 位。
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt 将识别十六进制格式的数字。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint 也可用。
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi 是基本的 base-10 int 解析的便利函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 解析函数在输入错误时返回错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}
