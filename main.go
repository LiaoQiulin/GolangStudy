package main

import (
	"fmt"
	"os"
)

// Go 为 printf 函数中的字符串格式提供了出色的支持。以下是一些常见字符串格式化任务的示例

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	// Go 提供了几个用于格式化一般 Go 值的打印“动词”。例如，这打印了我们的 point 结构的一个实例。
	fmt.Printf("struct1: %v\n", p)

	// 如果该值是一个结构，则 %+v 变体将包含结构的字段名称。
	fmt.Printf("struct2: %+v\n", p)

	// %#v 变体打印该值的 Go 语法表示，即会产生该值的源代码片段。
	fmt.Printf("struct3: %#v\n", p)

	// 要打印值的类型，请使用 %T。
	fmt.Printf("type: %T\n", p)

	// 格式化布尔值是直截了当的。
	fmt.Printf("bool: %t\n", true)

	// 格式化整数有很多选项。使用 %d 进行标准的 base-10 格式。
	fmt.Printf("int: %d\n", 123)

	// 这将打印二进制表示。
	fmt.Printf("bin: %b\n", 14)

	// 这将打印与给定整数对应的字符。
	fmt.Printf("char: %c\n", 33)

	// % x 提供十六进制编码。
	fmt.Printf("hex: %x\n", 456)

	// 浮点数还有几个格式化选项。对于基本的十进制格式，使用 %f。
	fmt.Printf("float1: %f\n", 78.9)

	// %e 和 %E 以科学记数法（略有不同的版本）格式化浮点数。
	fmt.Printf("float2: %e\n", 123400000.0)
	fmt.Printf("float3: %E\n", 123400000.0)

	// 对于基本的字符串打印，使用 %s。
	fmt.Printf("str1: %s\n", "\"string\"")

	// 要在 Go 源代码中使用双引号，请使用 %q。
	fmt.Printf("str2: %q\n", "\"string\"")

	// 与前面看到的整数一样，%x 以 base-16 呈现字符串，每个输入字节有两个输出字符。
	fmt.Printf("str3: %x\n", "hex this")

	// 要打印指针的表示，请使用 %p。
	fmt.Printf("pointer: %p\n", &p)

	// 格式化数字时，您通常希望控制结果图形的宽度和精度。要指定整数的宽度，请在动词的 % 之后使用数字。默认情况下，结果将右对齐并用空格填充。
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// 您还可以指定打印的浮点数的宽度，但通常您还希望使用 width.precision 语法同时限制小数精度。
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// 要左对齐，请使用 - 标志。
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// 您可能还想在格式化字符串时控制宽度，特别是要确保它们在类似表格的输出中对齐。对于基本的右对齐宽度。
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// 左对齐使用 - 标志与数字一样。
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// 到目前为止，我们已经看到了 Printf，它将格式化的字符串打印到 os.Stdout。
	// Sprintf 格式化并返回一个字符串，而不在任何地方打印它。
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// 您可以使用 Fprintf 格式化+打印到 os.Stdout 以外的 io.Writers。
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
