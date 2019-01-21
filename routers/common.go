package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/beego/i18n"
	"io"
	"os"
)

func DocsStatic(ctx *context.Context) {

	beego.Info("ctx:", ctx.Request.URL)
	if uri := ctx.Input.Param(":all"); len(uri) > 0 {
		lang := ctx.GetCookie("lang")
		if !i18n.IsExist(lang) {
			lang = "en-US"
		}
		beego.Info("lang:", lang)
		beego.Info("URL:", "docs/"  + "images/" + uri)
		f, err := os.Open("docs/"  + "images/" + uri)
		if err != nil {
			ctx.WriteString(err.Error())
			return
		}
		defer f.Close()

		_, err = io.Copy(ctx.ResponseWriter, f)
		if err != nil {
			ctx.WriteString(err.Error())
			return
		}
	}
}
