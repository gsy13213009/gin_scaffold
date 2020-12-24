package controller

import (
	"encoding/json"
	"fmt"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/dao"
	"github.com/gsy13213009/gin_scaffold/dto"
	"github.com/gsy13213009/gin_scaffold/middleware"
	"github.com/gsy13213009/gin_scaffold/public"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/info", adminLogin.AdminInfo)
	group.POST("/change_pwd", adminLogin.ChangePwd)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 获取管理员信息
// @Tags 管理员信息
// @ID /admin/info
// @Accept  json
// @Produce  json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutPut} "success"
// @Router /admin/info [get]
func (adminLogin *AdminController) AdminInfo(c *gin.Context) {
	// 1. 读取redis，转为struct
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	sessInfoStr := sessInfo.(string)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(sessInfoStr), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	// 2. 取出数据，封装成输出
	out := &dto.AdminInfoOutPut{
		ID:           adminSessionInfo.ID,
		LoginTime:    adminSessionInfo.LoginTime,
		Name:         adminSessionInfo.UserName,
		Avatar:       "",
		Introduction: "",
		Roles:        []string{"admin"},
	}
	middleware.ResponseSuccess(c, out)
}

// ChangePwd godoc
// @Summary 修改密码
// @Description 修改密码
// @Tags 管理员信息
// @ID /admin/change_pwd
// @Accept  json
// @Produce  json
// @Param body body dto.ChangePwdInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /admin/change_pwd [post]
func (adminLogin *AdminController) ChangePwd(c *gin.Context) {
	param := &dto.ChangePwdInput{}
	if err := param.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// 1. session 读取用户登录信息
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)), adminSessionInfo); err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2003, err)
		return
	}
	adminInfo := &dao.Admin{}
	// 2. sessionInfo.ID 读取数据库信息
	adminInfo, err = adminInfo.Find(c, tx, &dao.Admin{
		UserName: adminSessionInfo.UserName,
		Id:       adminSessionInfo.ID,
	})
	if err != nil {
		middleware.ResponseError(c, 2004, err)
		return
	}
	// 3. param.password + adminInfo.salt sha256 生成密码
	adminInfo.Password = public.GenSaltPassword(adminInfo.Salt, param.Password)
	// 4. 保存密码

	if err := adminInfo.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2005, err)
		return
	}
	middleware.ResponseSuccess(c, "")
}
