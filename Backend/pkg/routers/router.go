package routers

import (
	"net/http"
	"restaurant-system/pkg/controllers"
	"restaurant-system/pkg/middleware"
)

func RootRoutes(router *http.ServeMux, usersCon controllers.UserController, productsCon controllers.ProductController ) {
	mngr := &middleware.Manager{}
	mngr.Use(middleware.CorsMiddleware)

	usersRoutes(router, usersCon, mngr)
	productRoutes(router, productsCon, mngr )
}
