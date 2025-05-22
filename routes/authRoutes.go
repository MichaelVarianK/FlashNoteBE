package routes

import (
	atom_auth "github.com/KelXR/FlashNoteBE/atom/auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/api/auth")
	{
		// Regular Auth
		auth.POST("/register", atom_auth.CreateUser)
		auth.POST("/login", atom_auth.LoginUser)
		auth.POST("/refresh", atom_auth.RefreshAuth)

		// Google OAuth
		auth.GET("/google/login", atom_auth.GoogleLogin)
		auth.GET("/google/callback", atom_auth.GoogleCallback)
	}
}
