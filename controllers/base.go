package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	//添加校验登录过滤器
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("userName").(string)
		url := strings.Split(ctx.Request.RequestURI, "?")[0]
		if !ok && url != "/v1/user/login" {
			ctx.WriteString("请先登录！！！")
		}
	}
	beego.InsertFilter("/*",beego.BeforeRouter,FilterUser)
}

