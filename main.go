package main

import (
	"jyoyuu/config"
	"jyoyuu/controllers"
	"jyoyuu/middlewares"
	"jyoyuu/views"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
)

func main() {
	app := iris.New()
	//session中间件
	sess := sessions.New(sessions.Config{
		Cookie:       "_session_id",
		AllowReclaim: true,
	})
	app.Use(sess.Handler())
	app.Use(middlewares.Flash)
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
	app.Post("/login", controllers.LoginController)

	//读取配置文件
	serverConfig := new(config.ServerConfig)
	serverConfig.Read()
	app.Run(iris.Addr(serverConfig.ToString()), iris.WithCharset("UTF-8"))
}
