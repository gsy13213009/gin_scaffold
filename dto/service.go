package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/public"
)

type ServiceListInput struct {
	Info     string `json:"info" form:"info" comment:"关键词" example:"" validate:""`
	PageNo   int    `json:"page_no" form:"page_no" comment:"页数" example:"1" validate:"required"`
	PageSize int    `json:"page_size" form:"page_size" comment:"每页条数" example:"10" validate:"required"`
}

func (param *ServiceListInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceListOutPut struct {
	Total int64                   `json:"total" from:"total" comment:"总数" example:"" validate:""`
	List  []ServiceListItemOutPut `json:"list" from:"list" comment:"总数" example:"" validate:""`
}
type ServiceListItemOutPut struct {
	ID          int64  `json:"id" form:"id"`
	ServiceName string `json:"service_name" form:"service_name"`
	ServiceDesc string `json:"service_desc" form:"service_desc"`
	LoadType    int    `json:"load_type" form:"load_type"`
	ServiceAddr string `json:"service_addr" form:"service_addr"`
	Qps         int64  `json:"qps" form:"qps"`
	Qpd         int64  `json:"qpd" form:"qpd"`
	TotalNode   int    `json:"total_node" form:"qpd"`
}

type ServiceDeleteInput struct {
	ID     int64 `json:"id" form:"id" comment:"服务id" example:"20" validate:"required"`
}

func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
