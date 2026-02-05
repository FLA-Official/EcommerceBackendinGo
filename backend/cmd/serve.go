package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middlewares"
)

// Serve initializes dependencies and starts the HTTP server.
func Serve() {
	// load config
	cnf := config.GetConfig()
	// add config in a middleware constructor
	m := middlewares.NewMiddlewares(cnf)
	// Loading Dependencies
	productHandler := product.NewHandler(m)
	userHandler := user.NewHandler(m)
	// Dependency Injections
	server := rest.NewServer(cnf, productHandler, userHandler)
	// After All dependency Server starts
	server.Start()
}
