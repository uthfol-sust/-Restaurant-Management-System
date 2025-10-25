package cmd

import (
	"fmt"
	"net/http"
	"restaurant-system/pkg/config"
	"restaurant-system/pkg/connection"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/migrations"
	"restaurant-system/pkg/repositories"
	"restaurant-system/pkg/routers"
	"restaurant-system/pkg/services"

	"github.com/joho/godotenv"
)

func Serve() {
	godotenv.Load()
	config.SetConfig()

	db := connection.GetDB()
	migrations.Migrate(db)

	mux := http.NewServeMux()
	
	userRepo := repositories.NewUserRepository(db)
	productRepo := repositories.NewProductRepository(db)
	inventoryRepo := repositories.NewInventoryRepository(db)

	userService := services.NewUserService(userRepo)
	productService := services.NewProductService(productRepo)
	inventoryService := services.NewInventoryService(inventoryRepo)


	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
    inventoryController := controllers.NewInventoryController(inventoryService)

	

	routers.RootRoutes(mux,userController, productController, inventoryController)

	fmt.Println("Sever Runnig on Port 8080")
	http.ListenAndServe(":8080", mux)
}
