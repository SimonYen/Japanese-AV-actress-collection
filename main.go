package main

import (
	"jyoyuu/controllers"
	"jyoyuu/views"

	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()
	views.Init(true)
	app.RegisterView(views.Engine)
	app.HandleDir("/assets", iris.Dir("./views/static"))
	app.Get("/", controllers.HomeController)
	app.Run(iris.Addr(":9999"), iris.WithCharset("UTF-8"))
}
