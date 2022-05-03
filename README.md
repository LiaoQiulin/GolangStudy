# 切片
> 切片是 Go 中的关键数据类型，它为序列提供了比数组更强大的接口。

1. 与数组不同，切片仅按它们包含的元素（而不是元素的数量）进行类型化。要创建一个非零长度的空切片，请使用内置 make。在这里，我们制作了一段长度为 3 的字符串（最初为零值）。

```go
    s := make([]string, 3)
    fmt.Println("emp:", s)
```

2. 我们可以像使用数组一样设置和获取。
```go
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])
```

3. len 按预期返回切片的长度。
```go
    fmt.Println("len:", len(s))
```
4. 除了这些基本操作之外，切片还支持更多使其比数组更丰富的操作。一种是内置追加，它返回一个包含一个或多个新值的切片。请注意，我们需要接受 append 的返回值，因为我们可能会得到一个新的 slice 值。
```go
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)
```
5. 切片也可以复制。这里我们创建一个与 s 长度相同的空切片 c 并从 s 复制到 c 中。
```go
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)
```
6. 切片支持使用语法 slice[low:high] 的“切片”运算符。例如，这将获取元素 s[2]、s[3] 和 s[4] 的切片。
```go
    l := s[2:5]
    fmt.Println("sl1:", l)
```
7. 从 s[0] 切片到（但不包括）s[5]。
```go
    l = s[:5]
    fmt.Println("sl2:", l)
```
8. 从（并包括）s[2] 开始切片到（但不包括）s[5]。
```go
    l = s[:5]
    fmt.Println("sl2:", l)
```
9. 我们也可以在一行中为 slice 声明和初始化一个变量。
```go
    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
```
10. 切片可以组成多维数据结构。与多维数组不同，内部切片的长度可以变化。
```go
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
```
