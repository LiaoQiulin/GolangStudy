package main

import (
	"fmt"
	"time"
)

// Go 通过基于模式的布局支持时间格式化和解析。
func main() {
	p := fmt.Println

	// 这是一个根据 RFC3339 格式化时间的基本示例，使用相应的布局常量。
	t := time.Now()
	p(t.Format(time.RFC3339))

	// 时间解析使用与 Format 相同的布局值
	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	// Format 和 Parse 使用基于示例的布局。通常你会为这些布局使用 time 包里一个常量，但你也可以提供自定义布局。
	// 布局必须使用参考时间 Mon Jan 2 15:04:05 MST 2006 来显示格式化/解析给定时间/字符串的模式。
	// 示例时间必须完全如图所示：2006 代表年，小时为15，Mon代表星期几，等等。
	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p(t2)

	// 对于纯数字表示，您还可以将标准字符串格式与时间值的提取组件一起使用。
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse 将在格式错误的输入上返回错误，解释解析问题。
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
