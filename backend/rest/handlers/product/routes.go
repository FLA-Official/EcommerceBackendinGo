package product

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

// RegisterRoutes registers product-related routes on the provided mux.
func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
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
			http.HandlerFunc(h.GetProducts),
		),
	) // declaring Route

	mux.Handle("POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
		),
	) // declaring Route

	mux.Handle("GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductsByID),
		),
	) // declaring Route

	mux.Handle("PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProducts),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

	mux.Handle("DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProducts),
			h.middlewares.AuthenticateJWT,
		),
	) // declaring Route

}
