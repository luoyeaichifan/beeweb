// Copyright 2013 Beego Web authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package routers

import (
	"github.com/astaxie/beego"
	"github.com/luoyeaichifan/beeweb/models"
	"strings"
)

// DocsRouter serves about page.
type NavbarRouter struct {
	baseRouter
}

// Get implemented Get method for DocsRouter.
func (this *NavbarRouter) Get() {

	beego.Info("I am here")
	navbars := strings.Split(beego.AppConfig.String("navbar::types"), "|")
	this.Data["Pages"] = navbars
	url := this.Ctx.Request.URL.String()
	beego.Info("url:", url)
	for index, v := range navbars{
		if strings.Contains(url, "/"+ v) {

			this.Data[v] = true
			this.TplName = "navbar.html"
			beego.Info("v:", v)
			dRoot := models.GetDocByLocale(v)

			if dRoot == nil {
				this.Abort("404")
				return
			}

			link := this.GetString(":splat")

			beego.Info("link:", link)
			link = strings.TrimPrefix(link, v + "/")
			beego.Info("link:", link)
			var doc *models.DocNode
			if len(link) == 0 {
				if dRoot.Doc != nil {
					doc = dRoot.Doc
				} else {
					// todo
					beego.Info("I am here")
					//this.Redirect("/docs/beego/intro/", 302)
					this.Abort("404")
					return
				}
			} else {

				beego.Info("link:", link)
				beego.Info("1")
				doc, _ = dRoot.GetNodeByLink(link)
				if doc == nil {
					beego.Info("2")
					doc, _ = dRoot.GetNodeByLink(link + "/")
				}
			}



			if doc == nil {
				beego.Info("404")
				this.Abort("404")
				return
			}
			beego.Info("NavbarType:", v)
			this.Data["NavbarIndex"] = index
			this.Data["NavbarType"] = v
			this.Data["DocRoot"] = dRoot
			this.Data["Doc"] = doc
			this.Data["Title"] = doc.Name
			this.Data["Data"] = doc.GetContent()

			//beego.Info(fmt.Printf("doc:%#v", doc))
			//
			//for _,v1 := range doc.Docs {
			//	beego.Info(fmt.Printf("v1:%#v", v1))
			//	for _,v2 := range v1.Docs {
			//		beego.Info(fmt.Printf("v2:%#v", v2))
			//
			//		for _,v3 := range v2.Docs {
			//			beego.Info(fmt.Printf("v3:%#v", v3))
			//		}
			//	}
			//}
			return
		}
	}

	this.Abort("404")
	return

}


