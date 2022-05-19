package main

import (
	"fmt"
	s "strings"
)

// 标准库的字符串包提供了许多有用的字符串相关函数。以下是一些示例，可让您了解包装。

// 我们将 fmt.Println 别名为一个较短的名称，因为我们将在下面大量使用它。
var p = fmt.Println

func main() {

	// 这是字符串中可用函数的示例。
	// 由于这些是包中的函数，而不是字符串对象本身的方法，因此我们需要将要处理的字符串作为第一个参数传递给函数。
	// 您可以在字符串包文档中找到更多功能。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
}
