package controllers

import (
	"github.com/kataras/iris/v12"
)

func HomeController(ctx iris.Context) {
	ctx.View("home.html")
}
