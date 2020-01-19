package a_user

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"go-web-admin/app/service/s_user"
	"go-web-admin/library/e"
	"go-web-admin/library/inject"
	"go-web-admin/library/jwt"
	"go-web-admin/library/response"
	"go-web-admin/library/util"
	"net/http"
)

// 用户API管理对象
type Controller struct{}

type UserRequest struct {
	Username string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
	Role     []int  `v:"required-with#权限组 必须为 [] 列表"`
}

type SignInRequest struct {
	Username string `v:"required#账号不能为空"`
	Password string `v:"required#密码不能为空"`
}

// 用户登录接口
func Login(r *ghttp.Request) {

	data, _ := r.GetJson()

	var SignInData *SignInRequest
	if err := r.Parse(&SignInData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	authService := s_user.User{Username: data.GetString("username"), Password: data.GetString("password")}
	_, err := authService.Check()

	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_NOT_EXIST, err)
	} else {
		token, _ := jwt.GenerateToken(data.GetString("username"))
		data := map[string]string{
			"token": token,
		}
		response.Json(r, http.StatusOK, e.SUCCESS, data)
	}
}

func UserInfo(r *ghttp.Request) {

	data := map[string]interface{}{
		"name":    "admin",
		"user_id": "1",
		"access":  []string{"admin"},
		"token":   "token",
		"avatar":  "https://file.iviewui.com/dist/a0e88e83800f138b94d2414621bd9704.png",
	}
	response.Json(r, http.StatusOK, e.SUCCESS, data)
}

// RESTFul - GET
func (c *Controller) Get(r *ghttp.Request) {
	userService := s_user.User{
		Username: r.GetString("username"),
		PageNum:  util.GetPage(r),
		PageSize: g.Config().GetInt("setting.PageSize"),
	}
	id := r.GetInt("id")

	if id != 0 {
		userService.ID = id
		user, err := userService.Get()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_NOT_EXIST, "")
			return
		}
		user.Password = ""
		data := make(map[string]interface{})
		data["lists"] = user

		response.Json(r, http.StatusOK, e.SUCCESS, data)

	} else {
		total, err := userService.Count()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		users, err := userService.GetAll()
		if err != nil {
			response.Json(r, http.StatusBadRequest, e.ERROR_USER_GET_S_FAIL, "")
			return
		}
		for _, v := range users {
			v.Password = ""
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

	var UserData *UserRequest
	if err := r.Parse(&UserData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	userService := s_user.User{
		Username: data.GetString("username"),
		Password: data.GetString("password"),
		Role:     data.GetInts("role"),
	}

	if id, err := userService.Add(); err != e.SUCCESS {
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

	var UserData *UserRequest
	if err := r.Parse(&UserData); err != nil {
		response.Json(r, http.StatusBadRequest, e.INVALID_PARAMS, err.Error())
	}

	userService := s_user.User{
		ID:       r.GetInt("id"),
		Username: data.GetString("username"),
		Password: data.GetString("password"),
		Role:     data.GetInts("role"),
	}

	if id, err := userService.Edit(); err != e.SUCCESS {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_EDIT_FAIL, "")

	} else {
		err := inject.Obj.Common.UserAPI.LoadPolicy(id)
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

	userService := s_user.User{ID: r.GetInt("id")}
	_, err := userService.ExistByID()
	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_DELETE_FAIL, "")
		r.ExitAll()
	}
	user, err := userService.Get()

	err = userService.Delete()

	if err != nil {
		response.Json(r, http.StatusBadRequest, e.ERROR_USER_DELETE_FAIL, "")
		r.ExitAll()
	} else {
		_, _ = inject.Obj.Enforcer.DeleteUser(user.Username)
		response.Json(r, http.StatusOK, e.SUCCESS, nil)
	}
}
