package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/dto"
	"github.com/gsy13213009/gin_scaffold/middleware"
)

type ServiceController struct {
}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("service_list", service.ServiceList)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept  json
// @Produce  json
// @Param info query string true "关键词"
// @Param page_size query int true "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutPut} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	param := &dto.ServiceListInput{}
	if err := param.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	out := &dto.ServiceListOutPut{}
	middleware.ResponseSuccess(c, out)
}
