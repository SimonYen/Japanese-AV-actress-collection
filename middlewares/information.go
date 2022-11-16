package middlewares

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

/*
这个中间件的功能主要是从session中获取授权信息
*/
func Info(ctx iris.Context) {
	s := sessions.Get(ctx)
	status, _ := s.GetBoolean("authenticated")
	name := s.GetString("name")
	ctx.ViewData("isAuth", status)
	ctx.ViewData("name", name)
	ctx.Next()
}
