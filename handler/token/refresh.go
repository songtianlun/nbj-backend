package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
	"minepin-backend/model"
	"minepin-backend/pkg/errno"
	"minepin-backend/pkg/token"
	"minepin-backend/utils"
	"strconv"
)

func RefreshToken(c *gin.Context) {
	uid, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrUrl, nil)
	}
	u, err := model.GetUserByID(uid)
	if err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
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
	handler.SendResponse(c, nil, model.Token{UserID: u.Id, AccessToken: t, RefreshToken: r})
}
