package main

import (
	"fmt"
	"net"
	"net/url"
)

// URL 提供了一种统一的方式来定位资源。这是在 Go 中解析 URL 的方法。
func main() {

	// 我们将解析这个示例 URL，其中包括 scheme, authentication info, host, port, path, query params, and query片段。
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// 解析 URL 并确保没有错误。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// 访问该 scheme 很简单。
	fmt.Println(u.Scheme)

	// 用户包含所有认证信息；在此调用用户名和密码以获取各个值。
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// Host 包含主机名和端口（如果存在）。使用 SplitHostPort 提取它们。
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	// 这里我们提取路径和#后面的片段。
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	// 要以 k=v 格式的字符串获取查询参数，请使用 RawQuery。
	// 您还可以将查询参数解析为地图。解析的查询参数映射是从字符串到字符串切片，因此如果您只想要第一个值，请索引到 [0]。
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])

}
