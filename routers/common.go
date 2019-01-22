package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/beego/i18n"
	"io"
	"os"
	"strings"
)

func DocsStatic(ctx *context.Context) {

	//
	//if !strings.Contains(ctx.Request.URL.String(), "images") {
	//	return
	//}
	beego.Info("ctx:", ctx.Request.URL)
	if uri := ctx.Input.Param(":all"); len(uri) > 0 {
		lang := ctx.GetCookie("lang")
		if !i18n.IsExist(lang) {
			lang = "en-US"
		}
		beego.Info("lang:", lang)
		beego.Info("URL:", "docs/"  + "images/" + uri)
		rowUrl := ctx.Request.URL.String()
		url := ""
		if strings.Contains(rowUrl, "/go/") {
			url = "docs/" + "go/"  + "images/" + uri
		} else if strings.Contains(rowUrl, "/docs/"){
			url = "docs/" + "zh-CN/" + "images/" + uri
		} else if strings.Contains(rowUrl, "/js/"){
			url = "docs/" + "js/"  + "images/" + uri
		} else if strings.Contains(rowUrl, "/graphite/"){
			url = "docs/" + "graphite/"  + "images/" + uri
		} else if strings.Contains(rowUrl, "/crawler/"){
			url = "docs/" + "crawler/"  + "images/" + uri
		} else if strings.Contains(rowUrl, "/beego/"){
			url = "docs/" + "beego/"  + "images/" + uri
		}
		beego.Info("url:", url)
		f, err := os.Open(url)
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
