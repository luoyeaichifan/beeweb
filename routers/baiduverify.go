package routers

import (
	"html/template"
	"log"
)

// HomeRouter serves home page.
type BaiduVerify struct {
	baseRouter
}

// Get implemented Get method for HomeRouter.
func (this *BaiduVerify) Get() {
	//this.Data["IsHome"] = true
	this.TplName = "baidu_verify_Hbvw7qAUmJ.html"

	t, err := template.ParseFiles("baidu_verify_Hbvw7qAUmJ.html")
	if err != nil {
		log.Println("err:",err)
		return
	}
	t.Execute(this.baseRouter.Controller.Ctx.ResponseWriter, nil)
}