package controller

import (
	"encoding/json"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/dto"
	"github.com/gsy13213009/gin_scaffold/middleware"
	"github.com/gsy13213009/gin_scaffold/public"
)

type AdminController struct {
}

func AdminRegister(group *gin.RouterGroup) {
	adminLogin := &AdminController{}
	group.GET("/info", adminLogin.AdminInfo)
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
