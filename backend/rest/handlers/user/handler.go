package user

import "ecommerce/rest/middlewares"

type Handler struct {
	middlewares *middlewares.Middlewares
}

func NewHandler(m *middlewares.Middlewares) *Handler {
	return &Handler{
		middlewares: m,
	}
}
