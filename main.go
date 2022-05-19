package main

import (
	"fmt"
	"sort"
)

// Go 的 sort 包实现了内置和用户定义类型的排序。我们将首先查看内置函数的排序。
func main() {

	// 排序方法特定于内置类型；这是字符串的示例。
	// 请注意，排序是就地的，因此它会更改给定的切片并且不会返回新切片。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// 排序整数的示例。
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 我们还可以使用 sort 来检查切片是否已经排序。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
