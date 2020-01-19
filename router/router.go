package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-web-admin/app/api/a_menu"
	"go-web-admin/app/api/a_role"
	"go-web-admin/app/api/a_user"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// 统一路由注册.
func init() {
	// 用户模块 路由注册 - 使用执行对象注册方式
	s := g.Server()
	//s.Use(jwt.JWT, // 验证 token 是否正确
	//	permission.CasbinMiddleware, MiddlewareCORS) // 权限验证

	s.BindHandler("/token", a_user.Login)
	s.BindHandler("/userInfo", a_user.UserInfo)

	s.BindObjectRest("/api/v1/users/*id", new(a_user.Controller))
	s.BindObjectRest("/api/v1/roles/*id", new(a_role.Controller))
	s.BindObjectRest("/api/v1/menus/*id", new(a_menu.Controller))
}
