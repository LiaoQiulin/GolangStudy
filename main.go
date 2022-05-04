package main

import "fmt"

type rect struct {
	width, height int
}

// 此 area 方法的接收器类型为 *rect。
func (r *rect) area() int {
	return r.width * r.height
}

// 可以为 指针或值接收类型 定义方法。这是一个值接收的示例。
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// 这里我们调用为我们的结构定义的 2 个方法。
	fmt.Println("area: ", r.area())
	fmt.Println("perim:", r.perim())

	// Go 自动处理方法调用的值和指针之间的转换。
	// 您可能希望使用指针接收器类型来避免复制方法调用或允许方法改变接收结构。
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())
}
