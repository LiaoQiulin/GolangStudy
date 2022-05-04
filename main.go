package main

import (
	"fmt"
	"math"
)

// 这是几何形状的基本接口。
type geometry interface {
	area() float64
	perim() float64
}

// 我们将在 rect 和 circle 类型上实现此接口。
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

// 在 Go 中实现一个接口，我们只需要实现接口中的所有方法。在这里，我们在 rect 上实现 geometry 。
func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// circle 实现 geometry
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果变量具有接口类型，那么我们可以调用命名接口中的方法。这是一个通用测量函数，利用它来处理任何几何图形。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// circle 和 rect 结构类型都实现了几何接口，因此我们可以使用这些结构的实例作为参数来 measure。
	measure(r)
	measure(c)
}
