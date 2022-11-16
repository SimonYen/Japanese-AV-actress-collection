package controllers

import (
	"jyoyuu/database"
	"jyoyuu/models"
	"jyoyuu/utils"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

/*
这个文件保存这处理用户注册登录之类有关权限的控制器
*/
func LoginController(ctx iris.Context) {
	phone := ctx.PostValue("phone")
	psw := ctx.PostValue("psw")
	//检查数据库是否有这个电话号码的用户
	user := new(models.User)
	has, _ := database.DB.Where("phone=?", phone).Get(user)
	s := sessions.Get(ctx)
	if has {
		//检查密码是否一致
		if utils.CheckPasswordHash(psw, user.Password) {
			s.SetFlash("smsg", "登录成功🍌")
			//将状态，id写入到session中
			s.Set("authenticated", true)
			s.Set("id", user.Id)
			s.Set("name", user.Name)
		} else {
			s.SetFlash("emsg", "登录失败❌，密码错误！")
		}
	} else {
		s.SetFlash("emsg", phone+"尚未注册")
	}
	ctx.Redirect("/")
}

func RegisterController(ctx iris.Context) {
	phone := ctx.PostValue("phone")
	psw := ctx.PostValue("psw")
	name := ctx.PostValue("name")
	user := new(models.User)
	s := sessions.Get(ctx)
	has, _ := database.DB.Where("phone=?", phone).Get(user)
	if has {
		s.SetFlash("emsg", phone+"已经被注册！")
	} else {
		user.Name = name
		user.Password, _ = utils.HashPassword(psw)
		user.Phone = phone
		_, err := database.DB.Insert(user)
		if err != nil {
			s.SetFlash("emsg", "注册失败！数据库错误！")
		} else {
			s.SetFlash("smsg", "注册成功🍌")
			//将状态，id写入到session中
			s.Set("authenticated", true)
			s.Set("id", user.Id)
			s.Set("name", user.Name)
		}
	}
	ctx.Redirect("/")
}

func LogoutController(ctx iris.Context) {
	s := sessions.Get(ctx)
	s.Delete("authenticated")
	s.Delete("id")
	s.Delete("name")
	s.SetFlash("smsg", "登出成功。")
	ctx.Redirect("/")
}
