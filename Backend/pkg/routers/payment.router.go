package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func PaymentRoutes(router *http.ServeMux, controller controllers.PaymentController, manager *middleware.Manager) {
	router.Handle("POST /payments",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.CreatePayment),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /payments",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetAllPayments),
			middleware.AuthMiddleware,
		))

	router.Handle("GET /payments/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.GetPaymentByID),
			middleware.AuthMiddleware,
		))

	router.Handle("PUT /payments/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.UpdatePayment),
			middleware.AuthMiddleware,
		))

	router.Handle("DELETE /payments/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(controller.DeletePayment),
			middleware.AuthMiddleware,
		))
}
