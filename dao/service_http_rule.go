package dao

import (
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/public"
)

type HttpRule struct {
	ID             int64  `json:"id" gorm:"primary_key"`
	ServiceId      int64  `json:"service_id" grom:"column:service_id" description:"服务id"`
	RuleType       int8   `json:"rule_type" gorm:"column:rule_type" description:"匹配类型 0=url前缀url_prefix 1=域名domain "`
	Rule           string `json:"rule" gorm:"column:rule" description:"type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedHttps      int8   `json:"need_https" gorm:"column:need_https" description:"支持https 1=支持'"`
	NeedStripUri   int8   `json:"need_strip_uri" gorm:"column:need_strip_uri" description:"启用strip_uri 1=启用'"`
	NeedWebsocket  int8   `json:"need_websocket" gorm:"column:need_websocket" description:"是否支持websocket 1=支持'"`
	UrlRewrite     string `json:"url_rewrite" gorm:"column:url_rewrite" description:"url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔"`
	HeaderTransfor string `json:"header_transfor" gorm:"column:header_transfor" description:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔"`
}

func (t *HttpRule) TableName() string {
	return "gateway_service_http_rule"
}

func (t *HttpRule) Find(c *gin.Context, tx *gorm.DB, search *HttpRule) (*HttpRule, error) {
	out := &HttpRule{}
	err := tx.SetCtx(public.GetGinTraceContext(c)).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *HttpRule) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.SetCtx(public.GetGinTraceContext(c)).Save(t).Error
}
