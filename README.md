# 函数:可变参数函数
> 可以使用任意数量的尾随参数调用可变参数函数。例如，fmt.Println 是一个常见的可变参数函数。

1. 这是一个将任意数量的整数作为参数的函数。
```go
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}
```

2. 可以以通常的方式使用单个参数调用可变参数函数。
```go
    sum(1, 2)
    sum(1, 2, 3)
```

3. 如果切片中已经有多个参数，请像这样使用 func(slice...) 将它们应用于可变参数函数。
```go
    nums := []int{1, 2, 3, 4}
    sum(nums...)
```