package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Go 的 math/rand 包提供伪随机数生成。
func main() {

	// 例如，rand.Intn 返回一个随机整数 n，0 <= n < 100。
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// rand.Float64 返回一个 float64 f, 0.0 <= f < 1.0。
	fmt.Println(rand.Float64())

	// 这可用于生成其他范围内的随机浮点数，例如 5.0 <= f' < 10.0。
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 默认数字生成器是确定性的，因此默认情况下每次都会生成相同的数字序列。
	// 为了产生不同的序列，给它一个改变的种子。请注意，这对于您打算保密的随机数使用是不安全的，请使用 crypto/rand。
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	// 调用生成的 rand.Rand 就像 rand 包中的函数一样。
	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果您使用相同的数字为源播种，它会产生相同的随机数序列。
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
}
