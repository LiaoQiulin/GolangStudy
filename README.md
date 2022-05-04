# 指针
> Go 字符串是只读的字节片。语言和标准库特别对待字符串 - 作为以 UTF-8 编码的文本容器。在其他语言中，字符串由“字符”组成。在 Go 中，字符的概念称为符文——它是一个表示 Unicode 代码点的整数。[这篇 Go 博客文章](https://go.dev/blog/strings)很好地介绍了该主题。

1. s 是一个字符串，它分配了一个文字值，表示泰语中的单词“hello”。 Go 字符串文字是 UTF-8 编码的文本
```go
    const s = "สวัสดี"
```

2. 由于字符串等价于 []byte，这将产生存储在其中的原始字节的长度。
```go
    fmt.Println("Len:", len(s))
```

3. 对字符串进行索引会在每个索引处生成原始字节值。此循环生成构成 s 中代码点的所有字节的十六进制值。
```go
    for i := 0; i < len(s); i++ {
        fmt.Printf("%x ", s[i])
    }
    fmt.Println()
```

4. 要计算一个字符串中有多少个符文，我们可以使用 utf8 包。请注意，RuneCountInString 的运行时间取决于字符串的大小，因为它必须按顺序解码每个 UTF-8 符文。一些泰语字符由多个 UTF-8 代码点表示，因此这个计数的结果可能会令人惊讶。
```go
fmt.Println("Rune count:", utf8.RuneCountInString(s))
```

5. 范围循环专门处理字符串并解码每个符文及其在字符串中的偏移量。
```go
    for idx, runeValue := range s {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
```

6. 我们可以通过显式使用 utf8.DecodeRuneInString 函数来实现相同的迭代。
```go
    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
        runeValue, width := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%#U starts at %d\n", runeValue, i)
        w = width
    }
```
