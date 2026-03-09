package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"go-user-system/controller"
	"go-user-system/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	api := r.Group("/api")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/users", controller.GetUsers)
		api.DELETE("/users/:id", controller.DeleteUser)
	}

	return r
}
