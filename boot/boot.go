package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	_ "go-web-admin/app/model"
	"go-web-admin/library/inject"
	_ "go-web-admin/library/inject"
	_ "go-web-admin/router"
)

// 用于应用初始化。
func init() {

	_ = g.View()
	c := g.Config()
	s := g.Server()

	// 模板引擎配置
	//_ = v.AddPath("template")
	//v.SetDelimiters("${", "}")

	// glog配置
	logpath := c.GetString("setting.logpath")
	_ = glog.SetPath(logpath)

	// Web Server配置
	//s.SetServerRoot("public")
	s.SetLogPath(logpath)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(8000)

	AppSetting.PageSize = c.GetInt("setting.PageSize")

	_ = inject.LoadCasbinPolicyData()
}
