package controller

import (
	"encoding/json"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/dao"
	"github.com/gsy13213009/gin_scaffold/dto"
	"github.com/gsy13213009/gin_scaffold/middleware"
	"github.com/gsy13213009/gin_scaffold/public"
	"time"
)

type AdminLoginController struct {
}

func AdminLoginRegister(group *gin.RouterGroup) {
	adminLogin := &AdminLoginController{}
	group.POST("/login", adminLogin.AdminLogin)
	group.GET("/logout", adminLogin.AdminLogout)
}

// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录描述
// @Tags 管理员接口
// @ID /admin/login
// @Accept  json
// @Produce  json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin/login [post]
func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context) {
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	admin := &dao.Admin{}
	tx, err := lib.GetGormPool("default") // 拿到default连接池
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	admin, err = admin.LoginCheck(c, tx, params)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	// 设置session
	sessInfo := &dto.AdminSessionInfo{
		ID:        admin.Id,
		UserName:  admin.UserName,
		LoginTime: time.Now(),
	}
	sessBts, err := json.Marshal(sessInfo)
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}

	sess := sessions.Default(c)
	sess.Set(public.AdminSessionInfoKey, string(sessBts))
	_ = sess.Save()

	out := &dto.AdminLoginOutput{Token: admin.UserName}
	middleware.ResponseSuccess(c, out)
}

// AdminLogout godoc
// @Summary 管理员登出
// @Description 管理员登出描述
// @Tags 管理员接口
// @ID /admin/logout
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{} "success"
// @Router /admin/logout [get]
func (adminLogin *AdminLoginController) AdminLogout(c *gin.Context) {
	sess := sessions.Default(c)
	sess.Delete(public.AdminSessionInfoKey)
	_ = sess.Save()
	middleware.ResponseSuccess(c, "")
}
