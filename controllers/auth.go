package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

/*
这个文件保存这处理用户注册登录之类有关权限的控制器
*/
func LoginController(ctx iris.Context) {
	phone := ctx.PostValue("phone")
	psw := ctx.PostValue("psw")
	s := sessions.Get(ctx)
	s.SetFlash("smsg", phone+psw)
	ctx.Redirect("/")
}
