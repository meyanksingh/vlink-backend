package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/meyanksingh/vlink-backend/internal/app/controllers"
	middleware "github.com/meyanksingh/vlink-backend/internal/middleware"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	Auth := incomingRoutes.Group("/auth")
	{
		Auth.POST("/register", controller.Register)
		Auth.POST("/login", controller.Login)
		Auth.GET("/", middleware.JWTAuthMiddleware(), controller.Home)
	}

}

func UserRoutes(incomingRoutes *gin.Engine) {
	User := incomingRoutes.Group("/user")
	User.Use(middleware.JWTAuthMiddleware())
	{
		User.POST("/request", controller.SendFriendRequest)
		User.POST("/accept", controller.AcceptFriendRequest)
		User.POST("/decline", controller.DeclineFriendRequest)
		User.GET("/", controller.ListFriends)
		User.GET("/requests", controller.ListFriendRequests)
		User.DELETE("/remove", controller.RemoveFriend)
	}
}
