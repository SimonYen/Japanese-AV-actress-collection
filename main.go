package main

import "github.com/kataras/iris/v12"

func main() {
	app := iris.New()
	engine := iris.HTML("./", ".html")
	app.RegisterView(engine)
	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("こんにちは！")
	})
	app.Get("/hello", func(ctx iris.Context) {
		ctx.ViewData("title", "你好")
		ctx.ViewData("content", "女优库")
		ctx.View("HelloWorld.html")
	})
	app.Run(iris.Addr(":9999"), iris.WithCharset("UTF-8"))
}
