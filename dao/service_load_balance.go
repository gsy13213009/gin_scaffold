package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/public"
	"strings"
)

type LoadBalance struct {
	ID                     int64  `json:"id" gorm:"primary_key"`
	ServiceId              int64  `json:"service_id" grom:"column:service_id" description:"服务id"`
	CheckMethod            int8   `json:"check_method" gorm:"column:check_method" description:"检查方法 0 = tcpchk, 检测端口是否握手成功"`
	CheckTimeout           int    `json:"check_timeout" gorm:"column:check_timeout" description:"check超时时间, 单位s"`
	CheckInterval          int    `json:"check_interval" gorm:"column:check_interval" description:"检查间隔, 单位s"`
	RoundType              int8   `json:"round_type" gorm:"column:round_type" description:"轮询方式 0 = random 1 = round-robin 2 = weight_round-robin 3 = ip_hash"`
	IpList                 string `json:"ip_list" gorm:"column:ip_list" description:"ip列表"`
	WeightList             string `json:"weight_list" gorm:"column:weight_list" description:"权重列表"`
	ForbidList             string `json:"forbid_list" gorm:"column:forbid_list" description:"禁用ip列表"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" gorm:"column:upstream_connect_timeout" description:"建立连接超时, 单位s"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" gorm:"column:upstream_header_timeout" description:"获取header超时, 单位s"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" gorm:"column:upstream_idle_timeout" description:"链接最大空闲时间, 单位s"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" gorm:"column:upstream_max_idle" description:"最大空闲链接数"`
}

func (t *LoadBalance) TableName() string {
	return "gateway_service_load_balance"
}

func (t *LoadBalance) Find(c *gin.Context, tx *gorm.DB, search *LoadBalance) (*LoadBalance, error) {
	out := &LoadBalance{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *LoadBalance) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}

func (t *LoadBalance) GetIPListByModel() []string {
	return strings.Split(t.IpList, ",")
}
