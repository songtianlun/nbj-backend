package user

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
	"minepin-backend/model"
	"minepin-backend/pkg/auth"
	"minepin-backend/pkg/errno"
	"minepin-backend/pkg/token"
)

func Login(c *gin.Context) {
	var u model.UserBind
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	d, err := model.GetUser(u.Username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if err := auth.Compare(d.Password, u.Password); err != nil {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, err := token.Sign(token.Context{
		UUID: d.UUID,
		Role: d.Role,
	}, "")
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}
	handler.SendResponse(c, nil, model.Token{Token: t})

}
