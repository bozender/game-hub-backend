package routes

import (
	"github.com/gin-gonic/gin"
	"game-hub-backend/controllers"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes for browsing games and submitting games
	router.GET("/games", controllers.GetGames)
	router.POST("/games", controllers.CreateGame)

	// User authentication routes (we'll add these later)
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}
