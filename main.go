package main

import (
	"fmt"
	"net/http"
	"time"
)

// 在前面的示例中，我们研究了如何设置一个简单的 HTTP 服务器。
// HTTP 服务器可用于演示 context.Context 用于控制取消的用法。
// Context 携带跨越 API 边界和 goroutines 的截止日期、取消信号和其他请求范围的值。
func hello(w http.ResponseWriter, req *http.Request) {

	// 一个 context.Context 由 net/http 机器为每个请求创建，并且可用于 Context() 方法。
	ctx := req.Context()
	fmt.Println(time.Now(), "server: hello handler started")
	defer func() { fmt.Println(time.Now(), "server: hello handler ended") }()

	// 在向客户端发送回复之前等待几秒钟。这可以模拟服务器正在做的一些工作。
	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println(time.Now(), "server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
		return
	}

	go func() {
		<-ctx.Done()
		fmt.Println(time.Now(), "Done")
	}()

}

func main() {

	// 和以前一样，我们在“/hello”路由上注册我们的处理程序，然后开始服务。
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
