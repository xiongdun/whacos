package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"log"
	"net/url"
)

func main() {

	s := "https://www.yinzhengjie.com:8888/web/yinzhengjie/index.html?name=尹正杰&x=true#第三章"

	u, e := url.Parse(s)

	if e != nil {
		log.Fatal(e)
	}

	fmt.Println("使用的协议是：", u.Scheme)
	fmt.Println("主机名称是：", u.Host)
	fmt.Println("端口号是：", u.Port())
	fmt.Println("当前路径是:", u.Path)
	fmt.Println("请求信息是：", u.RawQuery)
	fmt.Println("参数部分是：", u.User)
	fmt.Println("当前的锚点是：", u.Fragment)
}

func socketTest() {

	windows.Socket()
}
