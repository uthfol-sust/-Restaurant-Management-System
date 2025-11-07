package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func SupplierRoutes(router *http.ServeMux, controllers controllers.SupplierController, manager *middleware.Manager) {
	router.Handle("POST /suppliers",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.CreateSupplier),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /suppliers",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.GetAllSupplier),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /suppliers/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.GetBySupplierID),
			middleware.AuthMiddleware,
		))

	router.Handle("PUT /suppliers/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.UpdateSupplier),
			middleware.AuthMiddleware,
		))

	router.Handle("DELETE /suppliers/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controllers.DeleteSupplier),
			middleware.AuthMiddleware,
		))
}
