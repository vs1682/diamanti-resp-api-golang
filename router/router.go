package router

import (
	"github.com/gin-gonic/gin"

	"restapi/controller"
	"restapi/middleware/jwt"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/auth", controller.Authenticate)

	r.Use(jwt.JWT())
	{
		r.GET("/albums", controller.GetAlbums)
		r.GET("/album/:id", controller.GetAlbum)
		r.POST("/album", controller.PostAlbum)
		r.PUT("/album/:id", controller.UpdateAlbum)
		r.DELETE("/album/:id", controller.DeleteAlbum)
	}

	return r
}
