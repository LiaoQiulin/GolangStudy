package main

import (
	"bufio"
	"fmt"
	"net/http"
)

// Go 标准库在 net/http 包中提供了对 HTTP 客户端和服务器的出色支持。在本例中，我们将使用它来发出简单的 HTTP 请求。
func main() {

	// 向服务器发出 HTTP GET 请求。
	// http.Get 是创建 http.Client 对象并调用其 Get 方法的便捷快捷方式；它使用具有有用默认设置的 http.DefaultClient 对象。
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 打印 HTTP 响应状态。
	fmt.Println("Response status:", resp.Status)

	// 打印响应正文的前 5 行。
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
