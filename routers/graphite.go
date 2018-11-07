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
	"io"
	"os"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/beego/i18n"

	"github.com/luoyeaichifan/beeweb/models"
)

// DocsRouter serves about page.
type GraphiteRouter struct {
	baseRouter
}

// Get implemented Get method for DocsRouter.
func (this *GraphiteRouter) Get() {

	beego.Info("I am here")
	this.Data["IsGraphite"] = true
	this.TplName = "graphite.html"

	dRoot := models.GetDocByLocale("graphite")

	if dRoot == nil {
		this.Abort("404")
		return
	}

	link := this.GetString(":splat")
	beego.Info(link)
	var doc *models.DocNode
	if len(link) == 0 {
		if dRoot.Doc != nil {
			doc = dRoot.Doc
		} else {
			// todo
			beego.Info("I am here")
			this.Redirect("/docs/beego/intro/", 302)
			return
		}
	} else {
		doc, _ = dRoot.GetNodeByLink(link)
		if doc == nil {
			doc, _ = dRoot.GetNodeByLink(link + "/")
		}
	}

	if doc == nil {
		this.Abort("404")
		return
	}

	this.Data["DocRoot"] = dRoot
	this.Data["Doc"] = doc
	this.Data["Title"] = doc.Name
	this.Data["Data"] = doc.GetContent()
}

func GraphiteStatic(ctx *context.Context) {
	if uri := ctx.Input.Param(":all"); len(uri) > 0 {
		lang := ctx.GetCookie("lang")
		if !i18n.IsExist(lang) {
			lang = "en-US"
		}

		f, err := os.Open("docs/" + lang + "/" + "images/" + uri)
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
