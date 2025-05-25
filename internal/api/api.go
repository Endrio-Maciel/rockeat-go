package api

import (
	"github.com/endrio-maciel/rockeat-go/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router      *chi.Mux
	UserService services.UserService
}

// func (api *Api) handleCreateUser(w http.ResponseWriter, r *http.Request) {

// }
