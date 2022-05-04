# 函数:多个返回值
> Go 内置了对多个返回值的支持。此功能在惯用的 Go 中经常使用，例如从函数返回结果值和错误值。

1. 此函数签名中的 (int, int) 表明该函数返回 2 个整数。
```go
func vals() (int, int) {
    return 3, 7
}
```

2. 使用来自多次赋值的调用的 2 个不同的返回值。
```go
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)
```

3. 如果您只需要返回值的子集，请使用空白标识符 _。
```go
    _, c := vals()
    fmt.Println(c)
```