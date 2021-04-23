package dto

import (
	"time"

	"github.com/LotteWong/giotto-gateway-admin/utils"
	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Username string `json:"username" form:"username" comment:"登录名称" example:"admin" validate:"required"`  // 登录名称
	Password string `json:"password" form:"password" comment:"登录密码" example:"123456" validate:"required"` // 登录密码
}

type LoginRes struct {
	Token string `json:"token" form:"token" comment:"登录令牌" example:"token" validate:""` // 登录令牌
}

type LoginSession struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	LoginAt  time.Time `json:"login_at"`
}

func (params *LoginReq) BindAndValidLoginReq(ctx *gin.Context) error {
	return utils.DefaultGetValidParams(ctx, params)
}
