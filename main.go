package main

import (
	"fmt"
	"time"
)

// Go 为时间和持续时间提供广泛的支持；这里有些例子。
func main() {
	p := fmt.Println

	// 我们将从获取当前时间开始。
	now := time.Now()
	p(now)

	// 您可以通过提供年、月、日等来构建时间结构。时间总是与位置相关联，即时区。
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 您可以按预期提取时间值的各个分量。
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 周一至周日工作日也可用。
	p(then.Weekday())

	// 这些方法比较两个时间，分别测试第一次发生在第二次之前、之后还是同时发生。
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub 方法返回一个 Duration 表示两次之间的间隔。
	diff := now.Sub(then)
	p(diff)

	// 我们可以以各种单位计算持续时间的长度。
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 您可以使用 Add 将时间提前一个给定的持续时间，或者使用 负号 将时间向后移动一个持续时间。
	p(then.Add(diff))
	p(then.Add(-diff))
}
