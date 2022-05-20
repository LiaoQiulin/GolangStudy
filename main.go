package main

import (
	"fmt"
	"time"
)

// 程序中的一个常见要求是获取自 Unix 纪元以来的秒数、毫秒数或纳秒数。以下是如何在 Go 中执行此操作。
func main() {

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// 您还可以将自纪元以来的整数秒或纳秒转换为相应的时间。
	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}
