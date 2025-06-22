package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Recoverer, middleware.Logger, api.Sessions.LoadAndSave)

	// csrfMiddleware := csrf.Protect(
	// 	[]byte(os.Getenv("GOBID_CSRF_KEY")),
	// 	csrf.Secure(false), //Dev only
	// )

	// api.Router.Use(csrfMiddleware)

	api.Router.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			// r.Get("/csrftoken", api.HandleGetCSRFtoken)
			r.Route("/users", func(r chi.Router) {
				r.Post("/signup", api.HandleSignupUser)
				r.Post("/login", api.HandleLoginUser)

				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware)
					r.Post("/logout", api.HandleLogoutUser)
				})
			})

			r.Route("/products", func(r chi.Router) {
				r.Group(func(r chi.Router) {
					r.Use(api.AuthMiddleware) // middleware para esse grupo
					r.Post("/", api.handleCreateProduct)

					// r.Get("ws/subscribe/{product_id}", api.HandleSubscribeUserToAuction)
				})
			})
		})
	})

}
