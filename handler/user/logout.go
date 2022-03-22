package user

import (
	"github.com/gin-gonic/gin"
	"mingin/handler"
	"mingin/model"
	"mingin/pkg/errno"
	"mingin/pkg/token"
)

func Logout(c *gin.Context) {
	claim, err := token.ParseRequest(c)
	if err != nil || claim.RtID <= 0 {
		handler.SendResponse(c, err, nil)
		return
	}
	if err = model.LogoutRefreshTokenWithTokenID(claim.RtID); err != nil {
		handler.SendResponse(c, errno.ErrLogoutRToken, nil)
	}
	handler.SendResponse(c, nil, "Bye ~")

}
