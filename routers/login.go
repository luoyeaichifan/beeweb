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
	"fmt"
	"github.com/astaxie/beego"
)

// DocsRouter serves about page.
type LoginRouter struct {
	baseRouter
}

// Get implemented Get method for DocsRouter.
func (this *LoginRouter) Get() {

	this.TplName = "login.html"
	return
}

func (this *LoginRouter) Post() {

	beego.Info("heihei")
	//this.Abort("404")

	username := this.GetString("username")
	password := this.GetString("password")


	fmt.Println("username:", username)
	fmt.Println("password:", password)

	this.Data["msg"] = "用户名或密码错误"
	this.TplName = "login.html"
	return
}


