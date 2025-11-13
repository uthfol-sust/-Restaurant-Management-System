package cmd

import (
	"fmt"
	"net/http"
	"restaurant-system/pkg/config"
	"restaurant-system/pkg/connection"
	"restaurant-system/pkg/core"
	"restaurant-system/pkg/migrations"
	"restaurant-system/pkg/routers"

	"github.com/joho/godotenv"
)

func Serve() {
	godotenv.Load()
	config.SetConfig()

	db := connection.GetDB()
	migrations.Migrate(db)

	mux := http.NewServeMux()

	controllersRoute := core.InitAppControllers(db)
	routers.RootRoutes(mux, *controllersRoute)

	fmt.Println("Sever Runnig on Port 8080")
	http.ListenAndServe(":8080", mux)
}
