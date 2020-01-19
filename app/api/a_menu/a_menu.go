package a_menu

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-web-admin/app/service/s_menu"
	"go-web-admin/library/e"
	"go-web-admin/library/response"
	"go-web-admin/library/util"
	"net/http"
)

// 用户API管理对象
type Controller struct{}

type MenuRequest struct {
	Name   string `v:"required#单名称 不能为空"`
	Type   string `v:"required#类型 不能为空，只能为 菜单/目录/按钮"`
	Path   string `v:"required#路径不能为空"`
	Method string `v:"required#方法不能为空"`
}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	menuService := s_menu.Menu{
		Name:     r.GetString("rname"),
		PageNum:  util.GetPage(r),
		PageSize: g.Config().GetInt("setting.PageSize"),
	}
	id := r.GetInt("id")

	if id != 0 {
		menuService.ID = id
		user, err := menuService.Get()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_EXIST_FAIL, "")
			return
		}
		data := make(map[string]interface{})
		data["lists"] = user

		response.Json(r, http.StatusOK, e.SUCCESS, data)

	} else {

		total, err := menuService.Count()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		users, err := menuService.GetAll()
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

	var MenuData *MenuRequest
	if err := r.Parse(&MenuData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	menuService := s_menu.Menu{
		Name:   data.GetString("name"),
		Type:   data.GetString("type"),
		Path:   data.GetString("path"),
		Method: data.GetString("method"),
	}

	if err := menuService.Add(); err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_ADD_FAIL, "")
	} else {
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}

}

// RESTFul - Put
func (c *Controller) Put(r *ghttp.Request) {
	data, _ := r.GetJson()

	var MenuData *MenuRequest
	if err := r.Parse(&MenuData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	menuService := s_menu.Menu{
		ID:     r.GetInt("id"),
		Name:   data.GetString("name"),
		Path:   data.GetString("path"),
		Method: data.GetString("method"),
	}

	if err := menuService.Edit(); err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_EDIT_FAIL, "")
	} else {
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}
}

// RESTFul - DELETE
func (c *Controller) Delete(r *ghttp.Request) {

	menuService := s_menu.Menu{ID: r.GetInt("id")}
	_, err := menuService.ExistByID()
	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_DELETE_FAIL, "")
		r.ExitAll()
	}
	err = menuService.Delete()
	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_MENU_DELETE_FAIL, "")
		r.ExitAll()
	} else {
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}

}
