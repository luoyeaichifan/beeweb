package routers



// HomeRouter serves home page.
type BaiduVerify struct {
	baseRouter
}

// Get implemented Get method for HomeRouter.
func (this *BaiduVerify) Get() {
	//this.Data["IsHome"] = true
	this.TplName = "baidu_verify_Hbvw7qAUmJ.html"
}