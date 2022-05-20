package main

import (
	"crypto/sha256"
	"fmt"
)

// SHA256 哈希经常用于计算二进制或文本 blob 的短标识。
// 例如，TL​​S/SSL 证书使用 SHA256 来计算证书的签名。这是在 Go 中计算 SHA256 哈希的方法。
func main() {
	s := "sha256 this string"

	// 在这里，我们从一个新的哈希开始。
	h := sha256.New()

	// 写入需要字节。如果您有一个字符串 s，请使用 []byte(s) 将其强制转换为字节。
	h.Write([]byte(s))

	// 这将最终的哈希结果作为字节切片。 Sum 的参数可用于附加到现有字节切片：通常不需要它。
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
