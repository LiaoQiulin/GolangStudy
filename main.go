package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Go 提供了对正则表达式的内置支持。以下是 Go 中常见的正则表达式相关任务的一些示例。
func main() {

	// 这将测试模式是否与字符串匹配。
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// 上面我们直接使用了字符串模式，但对于其他正则表达式任务，您需要编译优化的正则表达式结构。
	r, _ := regexp.Compile("p([a-z]+)ch")

	// 这些结构上有许多方法可用。这是我们之前看到的匹配测试。
	fmt.Println(r.MatchString("peach"))

	// 这会找到正则表达式的匹配项。
	fmt.Println(r.FindString("peach punch"))

	// 这也会找到第一个匹配项，但返回匹配项的开始和结束索引，而不是匹配的文本。
	fmt.Println("idx:", r.FindStringIndex("peach punch"))

	// 子匹配变体包括有关整个模式匹配和这些匹配中的子匹配的信息。例如，这将返回 p([a-z]+)ch 和 ([a-z]+) 的信息。
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// 同样，这将返回有关匹配和子匹配索引的信息。
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// 这些函数的所有变体适用于输入中的所有匹配项，而不仅仅是第一个。例如查找正则表达式的所有匹配项。
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// 同样，这将返回有关匹配和子匹配索引的信息。
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// 提供一个非负整数作为这些函数的第二个参数将限制匹配的数量。
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// 我们上面的例子有字符串参数，并使用了像 MatchString 这样的名称。我们还可以提供 []byte 参数并从函数名中删除 String。
	fmt.Println(r.Match([]byte("peach")))

	// 使用正则表达式创建全局变量时，您可以使用 Compile 的 MustCompile 变体。 MustCompile 恐慌而不是返回错误，这使得使用全局变量更安全。
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r)

	// regexp 包也可用于将字符串子集替换为其他值。
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// Func 变体允许您使用给定函数转换匹配的文本。
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
