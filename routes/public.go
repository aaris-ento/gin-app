package routes

import (
	"gin-app/bootstrap"

	"github.com/gin-gonic/gin"
)

func RegisterPublicRoutes(r *gin.Engine, app *bootstrap.App) {
	r.POST("/register", app.UserHandler.Register)
	r.POST("/login", app.UserHandler.Login)
	r.GET("/verify-email", app.UserHandler.VerifyEmail)
	r.POST("/forgot-password", app.UserHandler.ForgotPassword)
	r.POST("/reset-password", app.UserHandler.ResetPassword)

	r.GET("/products/:id", app.ProductHandler.GetProduct)
	r.GET("/products", app.ProductHandler.GetProducts)
}
