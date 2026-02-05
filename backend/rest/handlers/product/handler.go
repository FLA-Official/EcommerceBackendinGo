package product

import "ecommerce/rest/middlewares"

// Handler handles product-related HTTP requests.
type Handler struct {
	middlewares *middlewares.Middlewares
}

// NewHandler constructs a new product Handler with the provided middlewares.
func NewHandler(middlewares *middlewares.Middlewares) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
