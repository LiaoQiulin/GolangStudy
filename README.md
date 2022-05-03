# Switch
> Switch 语句在许多分支中表达条件。

1. 您可以使用逗号分隔同一 case 语句中的多个表达式。在此示例中，我们也使用了可选的默认情况。
```
    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }
```
2. 不带表达式的 switch 是表达 if/else 逻辑的另一种方式。在这里，我们还展示了 case 表达式如何可以是非常数。
```
    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }
```
3. 类型 switch 比较类型而不是值。您可以使用它来发现接口值的类型。在此示例中，变量 t 将具有与其子句对应的类型。
```
    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
```
