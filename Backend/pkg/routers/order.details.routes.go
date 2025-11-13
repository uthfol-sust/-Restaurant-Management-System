package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func OrderDetailsRoutes(router *http.ServeMux, controller controllers.OrderDetailController, manager *middleware.Manager) {
	router.Handle("POST /orderdetails",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.CreateOrderDetail),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /orderdetails/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetByOrderID),
			middleware.AuthMiddleware,
		))

	router.Handle("DELETE /orderdetails/{order_detail_id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.DeleteOrderDetail),
			middleware.AuthMiddleware,
		))
}
