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
	ID int64 `json:"id" form:"id" comment:"服务id" example:"20" validate:"required"`
}

func (param *ServiceDeleteInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}

type ServiceAddInput struct {
	ServiceName            string `json:"service_name" form:"service_name" validate:"required,is_valid_service_name"`
	ServiceDesc            string `json:"service_desc" form:"service_desc" validate:"required,max=255,min=1"`
	RuleType               int    `json:"rule_type" form:"rule_type" example:"" validate:"" comment:"匹配类型 0=url前缀url_prefix 1=域名domain"`
	Rule                   string `json:"rule" form:"rule" example:"" validate:"required" comment:"type=domain表示域名，type=url_prefix时表示url前缀"`
	NeedHttps              int    `json:"need_https" form:"need_https" example:"" validate:"" comment:"支持https 1=支持'"`
	NeedStripUri           int    `json:"need_strip_uri" form:"need_strip_uri" example:"" validate:"" comment:"启用strip_uri 1=启用'"`
	NeedWebsocket          int    `json:"need_websocket" form:"need_websocket" example:"" validate:"" comment:"是否支持websocket 1=支持'"`
	UrlRewrite             string `json:"url_rewrite" form:"url_rewrite" example:"" validate:"" comment:"url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔'"`
	HeaderTransfor         string `json:"header_transfor" form:"header_transfor" example:"" validate:"" comment:"header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔'"`
	OpenAuth               string `json:"open_auth" form:"open_auth" example:"" validate:"" comment:"是否开启权限"`
	BlackList              string `json:"black_list" form:"black_list" example:"" validate:"" comment:"黑名单ip"`
	WhiteList              string `json:"white_list" form:"white_list" example:"" validate:"" comment:"白名单ip"`
	ClientipFlowLimit      int    `json:"clientip_flow_limit" form:"clientip_flow_limit" example:"" validate:"" comment:"客户端ip限流"`
	ServiceFlowLimit       int    `json:"service_flow_limit" form:"service_flow_limit" example:"" validate:"" comment:"服务端限流"`
	RoundType              int    `json:"round_type" form:"round_type" example:"" validate:"" comment:"轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash"`
	IpList                 string `json:"ip_list" form:"ip_list" example:"" validate:"required" comment:"ip列表"`
	WeightList             string `json:"weight_list" form:"weight_list" example:"" validate:"" comment:"BS权重列表"`
	ForbidList             string `json:"forbid_list" form:"forbid_list" example:"" validate:"" comment:"禁用ip列表"`
	UpstreamConnectTimeout int    `json:"upstream_connect_timeout" form:"upstream_connect_timeout" example:"" validate:"" comment:"建立连接超时, 单位s"`
	UpstreamHeaderTimeout  int    `json:"upstream_header_timeout" form:"upstream_header_timeout" example:"" validate:"" comment:"获取header超时, 单位s"`
	UpstreamIdleTimeout    int    `json:"upstream_idle_timeout" form:"upstream_idle_timeout" example:"" validate:"" comment:"链接最大空闲时间, 单位s"`
	UpstreamMaxIdle        int    `json:"upstream_max_idle" form:"upstream_max_idle" example:"" validate:"" comment:"最大空闲链接数"`
}

func (param *ServiceAddInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
