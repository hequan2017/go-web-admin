package a_role

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"go-web-admin/app/service/s_role"
	"go-web-admin/library/e"
	"go-web-admin/library/inject"
	"go-web-admin/library/response"
	"go-web-admin/library/util"
	"net/http"
)

// 权限组 API管理对象
type Controller struct{}

type RoleRequest struct {
	Name string `v:"required#名称不能为空"`
	Menu []int  `v:"required-with#权限组 必须为 [] 列表"`
}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	roleService := s_role.Role{
		Name:     r.GetString("name"),
		PageNum:  util.GetPage(r),
		PageSize: g.Config().GetInt("setting.PageSize"),
	}
	id := r.GetInt("id")

	if id != 0 {
		roleService.ID = id
		user, err := roleService.Get()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_EXIST_FAIL, "")
			return
		}
		data := make(map[string]interface{})
		data["lists"] = user

		response.Json(r, http.StatusOK, e.SUCCESS, data)

	} else {

		total, err := roleService.Count()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		users, err := roleService.GetAll()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}

		data := make(map[string]interface{})
		data["lists"] = users
		data["total"] = total

		response.Json(r, http.StatusOK, e.SUCCESS, data)
	}
}

// RESTFul - POST
func (c *Controller) Post(r *ghttp.Request) {

	data, _ := r.GetJson()

	var RoleData *RoleRequest
	if err := r.Parse(&RoleData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	roleService := s_role.Role{
		Name: data.GetString("name"),
		Menu: data.GetInts("menu"),
	}

	if id, err := roleService.Add(); err != e.SUCCESS {
		response.Json(r, http.StatusBadRequest, err, "")
	} else {
		err := inject.Obj.Common.UserAPI.LoadPolicy(id)
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_LOAD_CASBIN_FAIL, "")
			r.ExitAll()
		}
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}

}

// RESTFul - Put
func (c *Controller) Put(r *ghttp.Request) {
	data, _ := r.GetJson()

	var RoleData *RoleRequest
	if err := r.Parse(&RoleData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	roleService := s_role.Role{
		ID:   r.GetInt("id"),
		Name: data.GetString("name"),
		Menu: data.GetInts("menu"),
	}

	if id, err := roleService.Edit(); err != e.SUCCESS {
		response.Json(r, http.StatusBadRequest, e.ERROR_ROLE_EDIT_FAIL, "")

	} else {
		err := inject.Obj.Common.RoleAPI.LoadPolicy(id)
		if err != nil {
			glog.Error(err)
			response.Json(r, http.StatusBadRequest, e.ERROR_LOAD_CASBIN_FAIL, "")
			r.ExitAll()
		}
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {

	roleService := s_role.Role{ID: r.GetInt("id")}
	_, err := roleService.ExistByID()
	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_ROLE_DELETE_FAIL, "")
		r.ExitAll()
	}
	role, err := roleService.Get()
	err = roleService.Delete()

	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_ROLE_DELETE_FAIL, "")
		r.ExitAll()
	} else {
		_, _ = inject.Obj.Enforcer.DeleteUser(role.Name)
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}

}
