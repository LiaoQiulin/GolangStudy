package main

import (
	b64 "encoding/base64"
	"fmt"
)

// Go 提供了对 base64 编码/解码的内置支持。
func main() {

	// 这是我们将编码/解码的字符串
	data := "abc123!?$*&()'-=@~"

	// Go 支持标准和 URL 兼容的 base64。这是使用标准编码器进行编码的方法。编码器需要一个 []byte，因此我们将字符串转换为该类型。
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	// 解码可能会返回一个错误，如果您还不知道输入格式是否正确，您可以检查该错误。
	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	fmt.Println()

	// 这使用与 URL 兼容的 base64 格式进行编码/解码。
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))
}
