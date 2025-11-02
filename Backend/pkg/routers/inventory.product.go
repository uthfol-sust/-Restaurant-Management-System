package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func InventoryProductRoutes(router *http.ServeMux, controllers controllers.InventoryProductController, manager *middleware.Manager) {
	router.Handle("POST /inventoryProducts",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.CreateInventoryProduct),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /inventoryProducts",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.GetAllInventoryProduct),
			middleware.AuthMiddleware,
		))

	router.Handle("PUT /inventoryProducts/id",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.UpdateInventoryProduct),
			middleware.AuthMiddleware,
		))
	router.Handle("DELETE /inventoryProducts/{product_id}/{inventory_id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.DeleteInventoryProduct),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /inventoryProducts/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.GetInventoryProductByID),
			middleware.AuthMiddleware,
		))

}
