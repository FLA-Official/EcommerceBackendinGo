package user

import (
	"ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// declaring Route for users
	mux.Handle("POST /users",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)

	mux.Handle("POST /users/login",
		manager.With(
			http.HandlerFunc(h.Login),
		),
	)

}
