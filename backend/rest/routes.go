package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// ctrl := func(w http.ResponseWriter, r *http.Request) {
	// 	log.Println("test")
	// }
	// // example of setting up middle where first it will be warped with middleware.Logger then it will be warped with middleware.Middleware2.
	// // it actually looks like this middleware.Middleware2(middleware.Logger(http.HandleFunc(ctrl))) => this proccess is done on Manager.go
	// mux.Handle("GET /route",
	// 	manager.With(
	// 		http.HandlerFunc(ctrl),
	// 	),
	// )
	mux.Handle("GET /products",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	) // declaring Route

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(handlers.CreateProduct),
		),
	) // declaring Route

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.GetProductsByID),
		),
	) // declaring Route

	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.UpdateProducts),
			middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(handlers.DeleteProducts),
			middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(handlers.CreateUser),
		),
	) // declaring Route

	mux.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(handlers.Login),
		),
	)

}
