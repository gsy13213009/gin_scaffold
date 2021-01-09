package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/public"
)

type TcpRule struct {
	ID        int64 `json:"id" gorm:"primary_key"`
	ServiceId int64 `json:"service_id" grom:"column:service_id" description:"服务id"`
	Port      int   `json:"port" grom:"column:port" description:"端口号"`
}

func (t *TcpRule) TableName() string {
	return "gateway_service_tcp_rule"
}

func (t *TcpRule) Find(c *gin.Context, tx *gorm.DB, search *TcpRule) (*TcpRule, error) {
	out := &TcpRule{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *TcpRule) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}
