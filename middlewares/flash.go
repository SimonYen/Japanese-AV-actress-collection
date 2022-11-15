package middlewares

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

func Flash(ctx iris.Context) {
	s := sessions.Get(ctx)
	smsg := s.GetFlashString("smsg")
	emsg := s.GetFlashString("emsg")
	if smsg != "" {
		ctx.ViewData("smsg", smsg)
		//同时清零
		s.SetFlash("smsg", "")
	}
	if emsg != "" {
		ctx.ViewData("emsg", emsg)
		//同时清零
		s.SetFlash("emsg", "")
	}
	ctx.Next()
}
