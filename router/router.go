package router

import (
	"github.com/gin-gonic/gin"
	"minepin-backend/handler/sd"
	"minepin-backend/handler/user"
	"minepin-backend/router/middleware"
	"net/http"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello, this is MinePin backend.")
	})

	g.POST("/login", user.Login)
	g.POST("/register", user.Create)

	u := g.Group("/v1/user")
	u.Use(middleware.AuthMiddleware())
	{
		u.GET("", user.List)
	}

	gsd := g.Group("/sd")
	{
		gsd.GET("/health", sd.HealthCheck)
		gsd.GET("/disk", sd.DiskCheck)
		gsd.GET("/cpu", sd.CPUCheck)
		gsd.GET("/ram", sd.RAMCheck)
	}

	return g
}
