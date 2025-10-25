package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func inventoryRoutes(router *http.ServeMux, inventories controllers.InventoryController, manager *middleware.Manager) {
	router.Handle("POST /inventories", manager.MiddlewareChain(
		http.HandlerFunc(inventories.CreateInventory),
		middleware.AuthMiddleware,
	))

	router.Handle("GET /inventories", manager.MiddlewareChain(
		http.HandlerFunc(inventories.GetAllInventory),
		middleware.AuthMiddleware,
	))

	router.Handle("GET /inventories/{id}", manager.MiddlewareChain(
		http.HandlerFunc(inventories.GetInventoryByID),
		middleware.AuthMiddleware,
	))

	router.Handle("PUT /inventories/{id}", manager.MiddlewareChain(
		http.HandlerFunc(inventories.UpdateInventory),
		middleware.AuthMiddleware,
	))

	router.Handle("DELETE /inventories/{id}", manager.MiddlewareChain(
		http.HandlerFunc(inventories.DeleteInventory),
		middleware.AuthMiddleware,
	))
}
