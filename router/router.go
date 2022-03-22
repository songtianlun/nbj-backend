package router

import (
	"github.com/gin-gonic/gin"
	"mingin/handler/sd"
	"mingin/handler/token"
	"mingin/handler/user"
	"mingin/router/middleware"
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

	// system status description
	gsd := g.Group("/sd")
	{
		gsd.GET("/health", sd.HealthCheck)
		gsd.GET("/disk", sd.DiskCheck)
		gsd.GET("/cpu", sd.CPUCheck)
		gsd.GET("/ram", sd.RAMCheck)
	}

	// v1 api
	v1 := g.Group("/v1")

	// login/register
	v1.POST("/login", user.Login)
	v1.POST("/register", user.Create)

	// user refresh token
	v1ur := v1.Group("/u/:id/refresh")
	v1ur.Use(middleware.AuthRefreshTokenMiddleware())
	{
		v1ur.GET("", token.RefreshToken)
	}
	// Admin User
	v1ua := v1.Group("/u/a")
	v1ua.Use(middleware.AuthAdminMiddleware())
	{
		v1ua.GET("/list", user.List)
	}

	// Average User
	v1u := v1.Group("/u")
	v1u.Use(middleware.AuthMiddleware())
	{
		v1u.GET("/:id/pref", user.GetPreferences)
		v1u.PUT("/:id/pref", user.SetPreferences)
		v1u.PUT("/:id", user.PutUpdateUser)
		v1u.GET("/:id", user.GetUser)
		v1u.GET("/:id/logout", user.Logout)
	}

	return g
}
