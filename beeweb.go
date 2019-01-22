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

// An open source project for official documentation and blog website of beego app framework.
package main

import (
	"os"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"

	"github.com/luoyeaichifan/beeweb/models"
	"github.com/luoyeaichifan/beeweb/routers"
)

const (
	APP_VER = "1.0.0"
)

// We have to call a initialize function manully
// because we use `bee bale` to pack static resources
// and we cannot make sure that which init() execute first.
func initialize() {
	models.InitModels()

	routers.IsPro = beego.BConfig.RunMode == "prod"
	if routers.IsPro {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}

	routers.InitApp()
}

func main() {

	initialize()

	beego.Info(beego.BConfig.AppName, APP_VER)

	//beego.InsertFilter("/docs/images/:all", beego.BeforeRouter, routers.DocsStatic)


	beego.InsertFilter("*/images/:all", beego.BeforeRouter, routers.DocsStatic)
	//beego.InsertFilter("/*", beego.BeforeRouter, routers.DocsStatic)

	if !routers.IsPro {
		beego.SetStaticPath("/static_source", "static_source")
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	//beego.SetStaticPath("/products/images", "products/images/")

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})
	//beego.Router("/community", &routers.CommunityRouter{})
	//beego.Router("/quickstart", &routers.QuickStartRouter{})
	//beego.Router("/video", &routers.VideoRouter{})
	//beego.Router("/products", &routers.ProductsRouter{})
	//beego.Router("/team", &routers.PageRouter{})
	//beego.Router("/about", &routers.AboutRouter{})
	//beego.Router("/donate", &routers.DonateRouter{})
	beego.Router("/docs/", &routers.DocsRouter{})
	beego.Router("/docs/*", &routers.DocsRouter{})
	//beego.Router("/blog", &routers.BlogRouter{})
	//beego.Router("/blog/*", &routers.BlogRouter{})
	beego.Router("/go", &routers.GoRouter{})
	beego.Router("/go/*", &routers.GoRouter{})
	beego.Router("/js", &routers.JSRouter{})
	beego.Router("/js/*", &routers.JSRouter{})
	beego.Router("/other", &routers.OtherRouter{})
	beego.Router("/other/*", &routers.OtherRouter{})
	beego.Router("/graphite", &routers.GraphiteRouter{})
	beego.Router("/graphite/*", &routers.GraphiteRouter{})
	beego.Router("/crawler", &routers.CrawlerRouter{})
	beego.Router("/crawler/*", &routers.CrawlerRouter{})
	beego.Router("/navbar", &routers.NavbarRouter{})
	beego.Router("/navbar/*", &routers.NavbarRouter{})
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
