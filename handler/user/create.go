package user

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
	"minepin-backend/model"
	"minepin-backend/pkg/errno"
	"minepin-backend/pkg/logger"
)

func Create(c *gin.Context) {
	var r CreateReq
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Nickname: r.Nickname,
		Password: r.Password,
		Email:    r.Email,
		Phone:    r.Phone,
	}

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		logger.ErrorF("error for valid data - %s", err.Error())
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateRep{
		Nickname: r.Nickname,
		Email:    r.Email,
		Phone:    r.Phone,
	}

	handler.SendResponse(c, nil, rsp)

}
