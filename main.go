package main

import (
	"encoding/xml"
	"fmt"
)

// Go 通过 encoding.xml 包为 XML 和类似 XML 的格式提供内置支持。

// Plant 将映射到 XML。与 JSON 示例类似，字段标签包含编码器和解码器的指令。
// 这里我们使用 XML 包的一些特殊功能： XMLName 字段名称指示代表此结构的 XML 元素的名称； id,attr 表示 Id 字段是 XML 属性而不是嵌套元素。
type Plant struct {
	XMLName xml.Name `xml:"plant"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Origin  []string `xml:"origin"`
}

func (p Plant) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func main() {
	coffee := &Plant{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"}

	// 代表我们工厂的 XML；使用 MarshalIndent 生成更易于阅读的输出。
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// 要将通用 XML 标头添加到输出，请显式添加它。
	fmt.Println(xml.Header + string(out))

	// 使用 Unmarhshal 将带有 XML 的字节流解析为数据结构。
	// 如果 XML 格式错误或无法映射到 Plant，则会返回描述性错误。
	var p Plant
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p)

	tomato := &Plant{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"}

	// parent>child>plant 字段标签告诉编码器将所有 plants  嵌套在 <parent><child>...
	type Nesting struct {
		XMLName xml.Name `xml:"nesting"`
		Plants  []*Plant `xml:"parent>child>plant"`
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plant{coffee, tomato}

	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
