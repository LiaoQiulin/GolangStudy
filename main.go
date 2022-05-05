package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// container 嵌入 base。嵌入看起来像一个没有名称的字段。
type container struct {
	base
	str string
}

func main() {

	// 当使用文字创建结构时，我们必须显式初始化嵌入；这里嵌入类型用作字段名称。
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// 直接在 co 上访问 base 的字段，例如 co.num
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// 或者，我们可以使用嵌入的类型名称拼出完整路径。
	fmt.Println("also num:", co.base.num)

	// 由于容器内嵌了base，所以base的方法也变成了容器的方法。在这里，我们调用一个直接从 co 的 base 嵌入的方法。
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// 使用方法嵌入结构可用于将接口实现赋予其他结构。在这里，我们看到一个容器现在实现了描述器接口，因为它嵌入了基础。
	var d describer = co
	fmt.Println("describer:", d.describe())

}
