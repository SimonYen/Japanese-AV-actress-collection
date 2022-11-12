package main

import (
	"jyoyuu/controllers"
	"jyoyuu/views"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
)

func main() {
	app := iris.New()
	// recover 中间件从任何异常中恢复，如果有异常，则写入500状态码（服务器内部错误）。
	app.Use(recover.New())

	requestLogger := logger.New(logger.Config{
		// 是否开启状态码
		Status: true,
		// 是否记录远程IP地址
		IP: true,
		// 是否呈现HTTP谓词
		Method: true,
		// 是否记录请求路径
		Path: true,
		// 是否开启查询追加
		Query: true,
	})
	app.Use(requestLogger)

	//初始化模板引擎
	views.Init(true)
	app.RegisterView(views.Engine)

	//处理静态文件
	app.HandleDir("/assets", iris.Dir("./views/static"))

	//路由
	app.Get("/", controllers.HomeController)
	app.Get("/about", func(ctx iris.Context) {
		ctx.View("about.html")
	})
	app.Run(iris.Addr(":9999"), iris.WithCharset("UTF-8"))
}
