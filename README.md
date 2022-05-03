# If/Else
> 在 Go 中使用 if 和 else 进行分支是直截了当的。

1. 你可以有一个没有 else 的 if 语句。
```
    if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
    }
```
2. 语句可以在条件句之前；此语句中声明的任何变量在所有分支中都可用
```
    if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
```
> 请注意，Go 中的条件不需要括号，但大括号是必需的。