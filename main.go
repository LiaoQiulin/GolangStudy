package main

import (
	"fmt"
	"net/http"
)

// 使用 net/http 包编写一个基本的 HTTP 服务器很容易。

// net/http 服务器中的一个基本概念是处理程序。处理程序是实现 http.Handler 接口的对象。
// 编写处理程序的常用方法是在具有适当签名的函数上使用 http.HandlerFunc 适配器。
func hello(w http.ResponseWriter, req *http.Request) {

	// 用作处理程序的函数将 http.ResponseWriter 和 http.Request 作为参数。
	// 响应编写器用于填写 HTTP 响应。这里我们的简单响应只是“hello\n”。
	fmt.Fprintf(w, "hello\n")
}

// 这个处理程序通过读取所有 HTTP 请求标头并将它们回显到响应正文中来做一些更复杂的事情。
func headers(w http.ResponseWriter, req *http.Request) {
	fmt.Println("in handle")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// 我们使用 http.HandleFunc 便利函数在服务器路由上注册我们的处理程序。它在 net/http 包中设置默认路由器，并将函数作为参数。
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// 最后，我们使用端口和处理程序调用 ListenAndServe。 nil 告诉它使用我们刚刚设置的默认路由器。

	http.ListenAndServe(":8090", nil)
}
