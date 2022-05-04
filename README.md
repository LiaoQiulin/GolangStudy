# 指针
> Go 支持指针，允许您在程序中传递对值和记录的引用。

1. zeroval 将获得 ival 的副本，该副本与调用函数中的副本不同。
```go
    func zeroval(ival int) {
        ival = 0
    }
```

2. zeroptr 相反有一个 *int 参数，这意味着它需要一个 int 指针。然后，函数体中的 *iptr 代码将指针从其内存地址取消引用到该地址的当前值。将值分配给取消引用的指针会更改引用地址处的值。
```go
    func zeroptr(iptr *int) {
        *iptr = 0
    }
```

3. &i 语法给出了 i 的内存地址，即指向 i 的指针。
```go
    zeroptr(&i)
    fmt.Println("zeroptr:", i)
```
