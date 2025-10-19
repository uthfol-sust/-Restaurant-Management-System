package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func usersRoutes(router *http.ServeMux, userController controllers.UserController, manager *middleware.Manager) {

	router.Handle("POST /register",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.CreateUser),
		))

	router.Handle("POST /login",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.Login),
		))

	router.Handle("GET /users",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.GetAllUsers),
		))

	router.Handle("GET /users/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.GetUserById),
		))

	router.Handle("GET /users/email/{email}",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.GetUserByEmail),
		))

	router.Handle("PUT /users/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.UpdateUsers),
		))
		
	router.Handle("DELETE /users/{id}",
		manager.MiddlewareChain(
			http.HandlerFunc(userController.DeleteUsers),
		))
}
