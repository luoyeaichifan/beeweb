package routers

import (
	"github.com/luoyeaichifan/beeweb/models"
)

type ProductsRouter struct {
	baseRouter
}

func (this *ProductsRouter) Get() {
	this.TplName = "products.html"
	this.Data["IsProducts"] = true
	this.Data["Products"] = models.Products
}
