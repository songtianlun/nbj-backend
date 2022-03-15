package middleware

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
	"minepin-backend/model"
	"minepin-backend/pkg/errno"
	"minepin-backend/pkg/token"
)

// AuthMiddleware 验证 Token 是否有效
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			handler.SendResponse(c, err, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

// AuthAdminMiddleware 验证 Token 及管理员身份是否有效
func AuthAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := token.ParseRequest(c)
		if err != nil {
			handler.SendResponse(c, err, nil)
			c.Abort()
			return
		}
		if claims.URole < model.ADMIN {
			handler.SendResponse(c, errno.ErrRole, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}

// AuthRefreshTokenMiddleware 验证 RefreshToken，验后即焚
func AuthRefreshTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRefreshTokenRequest(c); err != nil {
			handler.SendResponse(c, err, nil)
			c.Abort()
			return
		}
		c.Next()
	}
}
