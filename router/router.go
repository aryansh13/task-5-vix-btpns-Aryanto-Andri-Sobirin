package router

import (
	"go_restapi_gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// User Routes
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", controllers.RegisterUser)
		userGroup.POST("/login", controllers.LoginUser)
		userGroup.GET("/:userId", controllers.GetUserByID)
		userGroup.PUT("/:userId", controllers.UpdateUserByID)
		userGroup.DELETE("/:userId", controllers.DeleteUserByID)
	}

	// Photo Routes
	photoGroup := r.Group("/photos")
	{
		photoGroup.POST("/", controllers.UploadPhoto)
		photoGroup.GET("/", controllers.GetPhotos)
		photoGroup.PUT("/:photoId", controllers.UpdatePhotoByID)
		photoGroup.DELETE("/:photoId", controllers.DeletePhotoByID)
	}

	return r
}
