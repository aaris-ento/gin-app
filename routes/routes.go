package routes

import (
	"gin-app/bootstrap"
	"gin-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *bootstrap.App) *gin.Engine {
	r := gin.Default()

	r.Use(middlewares.ErrorHandler())

	RegisterPublicRoutes(r, app)
	RegisterAuthRoutes(r, app)

	return r
}
