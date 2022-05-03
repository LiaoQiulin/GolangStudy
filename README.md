# Maps
> Maps 是 Go 的内置关联数据类型（在其他语言中有时称为哈希或字典）

1. 要创建一个空 `Map` ，请使用内置 `make`：`make(map[key-type]val-type)`。

```go
    m := make(map[string]int)
```

2. 使用典型的 name[key] = val 语法设置键/值对。
```go
    m["k1"] = 7
    m["k2"] = 13

    fmt.Println("map:", m)
```

3. 获取名称为 [key] 的键的值。
```go
    v1 := m["k1"]
    fmt.Println("v1: ", v1)
```

4. 内置 len 在 Map 上调用时返回键/值对的数量。
```go
    fmt.Println("len:", len(m))
```

5. 内置 delete 从 Map 中删除键/值对。
```go
    delete(m, "k2")
    fmt.Println("map:", m)
```

6. 从映射中获取值时可选的第二个返回值指示该键是否存在于映射中。这可用于消除缺失键和具有零值（如 0 或“”）的键之间的歧义。这里我们不需要值本身，所以我们用空白标识符 _ 忽略它。
```go
    _, prs := m["k2"]
    fmt.Println("prs:", prs)
```

7. 还可以使用此语法在同一行中声明和初始化新映射。
```go
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
```

> 请注意，当使用 fmt.Println 打印时，地图会以 map[k:v k:v] 的形式出现。