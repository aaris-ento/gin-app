package routes

import (
	"gin-app/bootstrap"
	"gin-app/middlewares"
	"gin-app/models"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, app *bootstrap.App) {
	auth := r.Group("/")
	auth.Use(middlewares.JWT())

	auth.GET("/users/:id", app.UserHandler.GetUser)

	auth.GET("/profile",
		middlewares.CheckRole(string(models.RoleCustomer), string(models.RoleAdmin)),
		app.UserHandler.GetProfile,
	)

	auth.POST("/products",
		middlewares.CheckRole(string(models.RoleAdmin)),
		app.ProductHandler.CreateProduct,
	)

	auth.PUT("/products/:id",
		middlewares.CheckRole(string(models.RoleAdmin)),
		app.ProductHandler.UpdateProduct,
	)
}
