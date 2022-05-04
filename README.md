# Structs/结构
> Go 的结构是字段的类型化集合。它们对于将数据组合在一起以形成记录很有用。

1. 此 person 结构类型具有 name 和 age 字段。
```go
type person struct {
    name string
    age  int
}
```

2. newPerson 构造一个具有给定名称的person结构。
```go
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
}
```

3. 此语法创建一个新结构。
```go
    fmt.Println(person{"Bob", 20})
```

4. 您可以在初始化结构时命名字段。
```go
    fmt.Println(person{name: "Alice", age: 30})
```

5. 省略的字段将为零值。
```go
    fmt.Println(person{name: "Fred"})
```

6. & 前缀产生一个指向结构的指针。
```go
    fmt.Println(&person{name: "Ann", age: 40})
```

7. 将新的结构创建封装在构造函数中是惯用的
```go
    fmt.Println(newPerson("Jon"))
```

8. 使用点访问结构字段。
```go
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)
```

9. 还可以将点与结构指针一起使用 - 指针会自动取消引用。
```go
    sp := &s
    fmt.Println(sp.age)
```

10. 结构是可变的。
```go
    sp.age = 51
    fmt.Println(sp.age)
```
