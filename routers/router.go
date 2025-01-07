// @APIVersion 1.0.0
// @Title 接口文档
// @Description 系统服务端接口文档
package routers

import (
	"park/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/park/v1",
		beego.NSNamespace("/websocket",
			beego.NSInclude(&controllers.WebSocketController{}),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(&controllers.UserController{}),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(&controllers.CommentController{}),
		),
	)
	beego.AddNamespace(ns)
}
