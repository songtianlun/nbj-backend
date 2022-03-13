package middleware

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler"
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
