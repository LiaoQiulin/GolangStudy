# 函数
> 函数是 Go 的核心。

1. 这是一个函数，它接受两个整数并将它们的总和作为整数返回。Go 需要显式返回，即它不会自动返回最后一个表达式的值。
```go
func plus(a int, b int) int {
    return a + b
}
```

2. 当您有多个相同类型的连续参数时，您可以省略类似类型参数的类型名称，直到声明类型的最后一个参数。
```go
func plusPlus(a, b, c int) int {
    return a + b + c
}
```

3. 正如你所期望的那样调用一个函数，使用 name(args)。
```go
    res := plus(1, 2)
    fmt.Println("1+2 =", res)
    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
```