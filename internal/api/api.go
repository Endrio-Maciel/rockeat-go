package api

import (
	"github.com/alexedwards/scs/v2"
	"github.com/endrio-maciel/rockeat-go.git/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router         *chi.Mux
	UserService    services.UserService
	ProductService services.ProductService
	Sessions       *scs.SessionManager
}
