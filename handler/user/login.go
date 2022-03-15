package user

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
	"minepin-backend/model"
	"minepin-backend/pkg/auth"
	"minepin-backend/pkg/errno"
	"minepin-backend/pkg/token"
	"minepin-backend/utils"
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

	if pass := auth.PasswordVerify(d.Password, u.Password); pass {
		handler.SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	t, r, err := token.Sign(token.Claims{
		URole: d.Role,
		UID:   d.Id,
		UName: d.Nickname,
		UAddr: utils.GetAddrFromContext(c),
	})
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	go func() {
		log := model.UserLoginLog{
			UserUID:      d.Id,
			UserName:     d.Nickname,
			AccessToken:  t,
			RefreshToken: r,
			UserAddr:     utils.GetAddrFromContext(c),
		}
		if err = log.Create(); err != nil {
			return
		}
	}()

	handler.SendResponse(c, nil, model.Token{UserID: d.Id, AccessToken: t, RefreshToken: r})
}
