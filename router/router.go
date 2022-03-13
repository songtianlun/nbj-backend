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

	// Say Hello
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello, this is MinePin backend.")
	})

	// login/register
	g.POST("/login", user.Login)
	g.POST("/register", user.Create)
	// v1 api
	v1 := g.Group("/v1")
	// v1 api for user
	v1u := v1.Group("/user")
	v1u.Use(middleware.AuthMiddleware())
	{
		v1u.GET("", user.List)
	}
	v1u.Use(middleware.AuthRefreshTokenMiddleware())
	{
		v1u.POST("/:id/refresh")
	}

	// system status description
	gsd := g.Group("/sd")
	{
		gsd.GET("/health", sd.HealthCheck)
		gsd.GET("/disk", sd.DiskCheck)
		gsd.GET("/cpu", sd.CPUCheck)
		gsd.GET("/ram", sd.RAMCheck)
	}

	return g
}
