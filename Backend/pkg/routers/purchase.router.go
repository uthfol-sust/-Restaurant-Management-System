package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func PurchaseRoutes(router *http.ServeMux, controller controllers.PurchaseController, manager *middleware.Manager) {
	router.Handle("POST /purchases",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.CreatePurchase),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /purchases",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetAllPurchases),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /purchases/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetPurchaseByID),
			middleware.AuthMiddleware,
		))

	router.Handle("PUT /purchases/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.UpdatePurchase),
			middleware.AuthMiddleware,
		))

	router.Handle("DELETE /purchases/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.DeletePurchase),
			middleware.AuthMiddleware,
		))
}
