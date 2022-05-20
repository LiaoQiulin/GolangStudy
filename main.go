package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Go 提供了对 JSON 编码和解码的内置支持，包括与内置和自定义数据类型之间的往来。

// 我们将在下面使用这两个结构来演示自定义类型的编码和解码。
type response1 struct {
	Page   int
	Fruits []string
}

// 只有导出的字段才会在 JSON 中编码/解码。字段必须以大写字母开头才能导出。
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {

	// 首先，我们将了解如何将基本数据类型编码为 JSON 字符串。以下是原子值的一些示例。
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// 这里有一些切片和映射，它们按照您的期望编码为 JSON 数组和对象。
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// JSON 包可以自动编码您的自定义数据类型。它只会在编码输出中包含导出的字段，并且默认使用这些名称作为 JSON 键。
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// 您可以在结构字段声明上使用标签来自定义编码的 JSON 键名。检查上面 response2 的定义以查看此类标记的示例。
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// 我们需要提供一个变量，JSON 包可以在其中放置解码后的数据。此 map[string]interface{} 将保存字符串到任意数据类型的映射。
	var dat map[string]interface{}

	// 这是实际的解码，以及相关错误的检查。
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// 为了使用解码映射中的值，我们需要将它们转换为适当的类型。例如，这里我们将 num 中的值转换为预期的 float64 类型。
	num := dat["num"].(float64)
	fmt.Println(num)

	// 访问嵌套数据需要一系列转换。
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// 我们还可以将 JSON 解码为自定义数据类型。
	// 这样做的好处是为我们的程序添加了额外的类型安全性，并在访问解码数据时消除了对类型断言的需要。
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// 在上面的示例中，我们总是使用字节和字符串作为标准输出中数据和 JSON 表示之间的中间体。
	// 我们还可以将 JSON 编码直接流式传输到 os.Writers，如 os.Stdout，甚至是 HTTP 响应体。
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
