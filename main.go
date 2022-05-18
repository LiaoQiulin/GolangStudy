package main

import (
	"errors"
	"fmt"
)

// 在 Go 中，通过显式的、单独的返回值来传达错误是惯用的。

// 按照惯例，errors 是最后一个返回值，并且具有 error 类型，一个内置接口。
func f1(arg int) (int, error) {
	if arg == 42 {

		// errors.New 使用给定的错误消息构造一个基本错误值。
		return -1, errors.New("can't work with 42")

	}

	// 错误位置的 nil 值表示没有错误。
	return arg + 3, nil
}

// 通过在自定义类型上实现 Error() 方法，可以将自定义类型用作 error 。这是上面示例的一个变体，它使用自定义类型来显式表示参数错误。
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// 在这种情况下，我们使用 &argError 语法来构建一个新结构，为两个字段 arg 和 prob 提供值。
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// 如果您想以编程方式使用自定义错误中的数据，则需要通过类型断言将错误作为自定义错误类型的实例。
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
