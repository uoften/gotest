package main

import (
	"time"
	"github.com/kataras/iris"
	"github.com/kataras/iris/cache"
)

var markdownContents = []byte(`## Hello Markdown

This is a sample of Markdown contents

Features`)

//不应在包含动态数据的处理程序上使用缓存。
//缓存是静态内容的一个好的和必须的功能，即“关于页面”或整个博客网站，静态网站非常适合。
func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	app.Get("/", cache.Handler(10*time.Second), writeMarkdown)
	// 将其内容保存在第一个请求中并提供服务而不是重新加载内容。
	// 10秒后，它将被清除并重置。
	app.Run(iris.Addr(":8080"))
}

func writeMarkdown(ctx iris.Context) {
	//点击浏览器的刷新按钮多次，你会每10秒钟只看一次这个println
	println("Handler executed. Content refreshed.")
	ctx.Markdown(markdownContents)
}
/* 请注意，`StaticWeb`默认使用浏览器的磁盘缓存
   因此，在任何StaticWeb调用之后注册缓存处理程序
   为了更快的解决方案，服务器不需要跟踪响应
*/