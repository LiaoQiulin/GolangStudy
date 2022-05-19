package main

import (
	"fmt"
	"sort"
)

// 有时我们会希望按照不同于自然顺序的方式对集合进行排序。
// 例如，假设我们想按字符串的长度而不是按字母顺序对字符串进行排序。这是 Go 中自定义排序的示例。

// 为了在 Go 中按自定义函数排序，我们需要相应的类型。
// 在这里，我们创建了一个 byLength 类型，它只是内置 []string 类型的别名。
type byLength []string

// 我们在我们的类型上实现了 sort.Interface - Len、Less 和 Swap，因此我们可以使用 sort 包的通用 Sort 函数。
// Len 和 Swap 通常在类型之间是相似的，而 Less 将保存实际的自定义排序逻辑。
// 在我们的例子中，我们希望按照字符串长度增加的顺序进行排序，所以我们在这里使用 len(s[i]) 和 len(s[j])
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// 有了所有这些，我们现在可以通过将原始 fruits 切片转换为 byLength 来实现自定义排序，然后在该类型切片上使用 sort.Sort。
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
