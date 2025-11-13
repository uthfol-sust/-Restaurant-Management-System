package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func OrderRoutes(router *http.ServeMux, controller controllers.OrderController, manager *middleware.Manager) {
	router.Handle("POST /orders",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.CreateOrder),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /orders",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetAllOrders),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /orders/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetOrderByID),
			middleware.AuthMiddleware,
		))

	router.Handle("PUT /orders/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.UpdateOrder),
			middleware.AuthMiddleware,
		))

	router.Handle("DELETE /orders/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.DeleteOrder),
			middleware.AuthMiddleware,
		))
}
