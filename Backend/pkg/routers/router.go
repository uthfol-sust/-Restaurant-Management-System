package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func RootRoutes(router *http.ServeMux, controllers controllers.UserController) {
	mngr := &middleware.Manager{}
	mngr.Use(middleware.CorsMiddleware)

	usersRoutes(router, controllers, mngr)
	productRoutes(router)
}
