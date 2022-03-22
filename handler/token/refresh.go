package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"mingin/handler"
	"mingin/model"
	"mingin/pkg/errno"
	"mingin/pkg/token"
	"mingin/utils"
	"strconv"
)

func RefreshToken(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrUrl, nil)
		return
	}
	u, err := model.GetUserByID(uid)
	if err != nil || u.Id != uid {
		handler.SendResponse(c, errno.ErrUID, nil)
		return
	}
	t, r, err := token.Sign(token.Claims{
		URole:          u.Role,
		UID:            u.Id,
		UName:          u.Nickname,
		UAddr:          utils.GetAddrFromContext(c),
		StandardClaims: jwt.StandardClaims{},
	})
	if err != nil {
		handler.SendResponse(c, errno.ErrToken, nil)
		return
	}

	go func() {
		log := model.UserLoginLog{
			UserUID:      uid,
			UserName:     u.Nickname,
			AccessToken:  t,
			RefreshToken: r,
			UserAddr:     utils.GetAddrFromContext(c),
			Type:         "rt",
		}
		if err = log.Create(); err != nil {
			return
		}
	}()

	handler.SendResponse(c, nil, model.Token{UserID: u.Id, Nickname: u.Nickname, AccessToken: t, RefreshToken: r})
	return
}
