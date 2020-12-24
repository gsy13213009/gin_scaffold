package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/gsy13213009/gin_scaffold/public"
	"time"
)

type AdminInfoOutPut struct {
	ID           int       `json:"id"`
	Name         string    `json:"user_name"`
	LoginTime    time.Time `json:"login_time"`
	Avatar       string    `json:"avatar"`
	Introduction string    `json:"introduction"`
	Roles        []string  `json"roles"`
}

type ChangePwdInput struct {
	Password string `json"passwrod" form:"password" comment:"密码" expample:"123456" validate:"required"`
}

func (param *ChangePwdInput) BindValidParam(c *gin.Context) error {
	return public.DefaultGetValidParams(c, param)
}
