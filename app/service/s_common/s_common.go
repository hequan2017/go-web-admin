package s_common

import (
	"go-web-admin/app/service/s_menu"
	"go-web-admin/app/service/s_role"
	"go-web-admin/app/service/s_user"
)

type Common struct {
	UserAPI *s_user.User `inject:""`
	RoleAPI *s_role.Role `inject:""`
	MenuAPI *s_menu.Menu `inject:""`
}
