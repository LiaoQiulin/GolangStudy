package main

import (
	"os"
	"text/template"
)

// Go 为创建动态内容或使用 text/template 包向用户显示自定义输出提供了内置支持。
// 名为 html/template 的同级包提供了相同的 API，但具有额外的安全功能，应该用于生成 HTML。

func main() {

	// 我们可以创建一个新模板并从字符串中解析其主体。模板是是静态文本和包含在{{...}}中用于动态插入内容的“动作”的混合体。
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// 或者，我们可以使用 template.Must 函数来恐慌(panic) Parse 返回错误。这对于在全局范围内初始化的模板特别有用。
	t1 = template.Must(t1.Parse("Value: {{.}}\n"))

	// 通过“执行”模板，我们为其操作生成具有特定值的文本。 {{.}} 动作被作为参数传递给 Execute 的值替换。
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"C++",
		"C#",
	})

	// 我们将在下面使用辅助函数。
	Create := func(name, t string) *template.Template {
		return template.Must(template.New(name).Parse(t))
	}

	// 如果数据是一个结构，我们可以使用 {{.FieldName}} 操作来访问它的字段。执行模板时，应导出字段以供访问。
	t2 := Create("t2", "Name: {{.Name}}\n")

	t2.Execute(os.Stdout, struct {
		Name string
	}{"Jane Doe"})

	// 这同样适用于map；
	t2.Execute(os.Stdout, map[string]string{
		"Name": "Mickey Mouse",
	})

	// if/else 为模板提供条件执行。如果一个值是类型的默认值，例如 0、空字符串、nil 指针等，则该值被认为是 false。
	// 此示例演示了模板的另一个功能：在操作中使用 - 来修剪空白。
	t3 := Create("t3",
		"{{if . -}} yes {{else -}} no {{end}}\n")
	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, " ")
	t3.Execute(os.Stdout, "")

	// range 块让我们可以遍历切片、数组、映射或通道。在范围块内 {{.}} 设置为迭代的当前项。
	t4 := Create("t4",
		"Range: {{range .}}{{.}} {{end}}\n")
	t4.Execute(os.Stdout,
		[]string{
			"Go",
			"Rust",
			"C++",
			"C#",
		})
}
