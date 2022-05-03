# For
> for 是 Go 唯一的循环结构。以下是 for 循环的一些基本类型。

1. 最基本的类型，具有单一条件。
```
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
```
2. 经典的 for 循环: 初始/条件/之后 。
```
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }
```
3. 没有条件的 for 将重复循环，直到您跳出循环或从封闭函数返回。
```
    for {
        fmt.Println("loop")
        break
    }
```
4. 您还可以继续循环的下一次迭代。
```
    for n := 0; n <= 5; n++ {
        if n%2 == 0 {
            continue
        }
        fmt.Println(n)
    }
```