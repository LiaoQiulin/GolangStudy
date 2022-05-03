# 数组
> 在 Go 中，数组是具有特定长度的元素的编号序列。

1. 在这里，我们创建了一个数组 a 将恰好容纳 5 个整数。元素的类型和长度都是数组类型的一部分。默认情况下，数组是零值的，对于整数意味着 0。

```go
    var a [5]int
    fmt.Println("emp:", a)
```

2. 我们可以使用 array[index] = value 语法在索引处设置一个值，并使用 array[index] 获取一个值
```go
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
```

3. 内置 len 返回数组的长度。
```go
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])
```
4. 此语法在一行中声明和初始化一个数组。
```go
    b := [5]int{1, 2, 3, 4, 5}
```
5. 数组类型是一维的，但您可以组合类型来构建多维数据结构。
```go
    var twoD [2][3]int
```