package router

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/verkatech/neko/api"
	"github.com/verkatech/neko/pkg/config"
)

func InitRoutes(g *gin.Engine) {

	g.Use(api.AuthMiddleware())

	// serve frontend static files
	g.Use(static.Serve("/app", static.LocalFile(config.Server.StaticPath, true)))

	apiRoute := g.Group("api/v1") // api version 1

	// user related endpoints
	apiRoute.POST("/users", api.Register)
	apiRoute.POST("/auth", api.Login)
}
