package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/public"
)

type AccessControl struct {
	ID                int64  `json:"id" gorm:"primary_key"`
	ServiceID         int64  `json:"service_id" grom:"column:service_id" description:"服务id"`
	OpenAuth          int    `json:"open_auth" gorm:"column:open_auth" description:"是否开启权限 1=开启"`
	BlackList         string `json:"black_list" gorm:"colunm:black_list" description:"黑名单ip"`
	WhiteList         string `json:"white_list" gorm:"colunm:white_list" description:"白名单ip"`
	WhiteHostName     string `json:"white_host_name" gorm:"colunm:white_host_name" description:"白名单主机"`
	ClientIPFlowLimit int64  `json:"client_flow_limit" gorm:"colunm:client_flow_limit" description:"客户端ip限流"`
	ServiceFlowLimit  int64  `json:"service_flow_limit" gorm:"colunm:service_flow_limit" description:"服务端限流"`
}

func (t *AccessControl) TableName() string {
	return "gateway_service_access_control"
}

func (t *AccessControl) Find(c *gin.Context, tx *gorm.DB, search *AccessControl) (*AccessControl, error) {
	out := &AccessControl{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *AccessControl) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}
