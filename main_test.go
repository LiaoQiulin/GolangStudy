package main

import (
	"fmt"
	"testing"
)

// 单元测试是编写有原则的 Go 程序的重要组成部分。
// 测试包提供了我们编写单元测试和 go test 命令运行测试所需的工具。

// 为了演示起见，此代码位于 main 包中，但它可以是任何包。
// 测试代码通常与它测试的代码位于同一个包中。

// 我们将测试这个整数最小值的简单实现。
// 通常，我们正在测试的代码将在一个名为 intutils.go 的源文件中，然后它的测试文件将被命名为 intutils_test.go。
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 测试是通过编写一个名称以 Test 开头的函数来创建的。
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		// t.Error* 将报告测试失败但继续执行测试。 t.Fatal* 将报告测试失败并立即停止测试。
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// 编写测试可能是重复的，因此使用表格驱动样式是惯用的，其中测试输入和预期输出列在表格中，并且单个循环遍历它们并执行测试逻辑。
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// t.Run 启用运行“子测试”，每个表条目一个。这些在执行 go test -v 时单独显示。
	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// 基准测试通常放在 _test.go 文件中，并以 Benchmark 开头。
// 测试运行器多次执行每个基准函数，每次运行增加 b.N 直到它收集到精确的测量值。
func BenchmarkIntMin(b *testing.B) {

	// 通常，基准测试运行一个我们在循环 b.N 次中进行基准测试的函数。
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

/*
$ go test -v
== RUN   TestIntMinBasic
--- PASS: TestIntMinBasic (0.00s)
=== RUN   TestIntMinTableDriven
=== RUN   TestIntMinTableDriven/0,1
=== RUN   TestIntMinTableDriven/1,0
=== RUN   TestIntMinTableDriven/2,-2
=== RUN   TestIntMinTableDriven/0,-1
=== RUN   TestIntMinTableDriven/-1,0
--- PASS: TestIntMinTableDriven (0.00s)
    --- PASS: TestIntMinTableDriven/0,1 (0.00s)
    --- PASS: TestIntMinTableDriven/1,0 (0.00s)
    --- PASS: TestIntMinTableDriven/2,-2 (0.00s)
    --- PASS: TestIntMinTableDriven/0,-1 (0.00s)
    --- PASS: TestIntMinTableDriven/-1,0 (0.00s)
PASS
ok      examples/testing-and-benchmarking    0.023s

$ go test -bench=.
goos: darwin
goarch: arm64
pkg: examples/testing
BenchmarkIntMin-8 1000000000 0.3136 ns/op
PASS
ok      examples/testing-and-benchmarking    0.351s

*/
