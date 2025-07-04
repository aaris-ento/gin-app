package bootstrap

import (
	"gin-app/handlers"
	"gin-app/repository"
	"gin-app/services"
)

type App struct {
	UserHandler    *handlers.UserHandler
	ProductHandler *handlers.ProductHandler
}

func Init() *App {
	db := InitDB()
	asyncClient := InitAsynqClient()

	taskService := services.NewTaskService(asyncClient)

	userRepo := repository.NewUserRepo(db)
	userService := services.NewUserService(userRepo, taskService)
	userHandler := handlers.NewUserHandler(userService)

	productRepo := repository.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	return &App{
		UserHandler:    userHandler,
		ProductHandler: productHandler,
	}
}
