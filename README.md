# 函数:递归
> Go 支持递归函数。

1. 这个 fact 函数调用自己，直到它到达 fact（0）的基本情况。
```go
func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
```

2. 闭包也可以是递归的，但这需要在定义闭包之前用类型化的 var 显式声明。
```go
func main() {
	fmt.Println(fact(7))

	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
```
