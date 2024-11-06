package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meyanksingh/vlink-backend/internal/app/controllers"
	middleware "github.com/meyanksingh/vlink-backend/internal/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup)
	incomingRoutes.POST("users/login", controller.Login)
	incomingRoutes.GET("users/home", middleware.JWTAuthMiddleware(), controller.Home)
}

func MainRoutes(incomiongRoutes *gin.Engine) {
	incomiongRoutes.GET("/", controller.HomePage)
}
