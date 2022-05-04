# 函数:闭包
> Go 支持匿名函数，可以形成闭包。当您想内联定义函数而不必命名时，匿名函数很有用。

1. 这个函数 intSeq 返回另一个函数，我们在 intSeq 的主体中匿名定义。返回的函数关闭变量 i 以形成一个闭包。
```go
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}
```

2. 我们调用 intSeq，将结果（一个函数）分配给 nextInt。这个函数值捕获自己的 i 值，每次调用 nextInt 时都会更新。
```go
    nextInt := intSeq()
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())
```

3. 要确认该状态对于该特定功能是唯一的，请创建并测试一个新状态。
```go
    newInts := intSeq()
    fmt.Println(newInts())
```