# Rang
> range 迭代各种数据结构中的元素。让我们看看如何将 range 与我们已经学过的一些数据结构一起使用。

1. 这里我们使用 range 对切片中的数字求和。数组也是这样工作的。
```go
 work like this too.

    nums := []int{2, 3, 4}
    sum := 0
    for _, num := range nums {
        sum += num
    }
    fmt.Println("sum:", sum)
```

2. 数组和切片上的 range 为每个条目提供索引和值。上面我们不需要索引，所以我们用空白标识符 _ 忽略它。有时我们实际上想要索引。
```go
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
```

3. Map 上的 range 迭代键/值对。
```go
    for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
```

4. range 也可以只遍历 Map 的键。
```go
    for k := range kvs {
        fmt.Println("key:", k)
    }
```

5. 字符串范围迭代 Unicode 代码点。第一个值是符文的起始字节索引，第二个值是符文本身。
```go
    for i, c := range "go" {
        fmt.Println(i, c)
    }
```