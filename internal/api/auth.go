package api

import (
	"net/http"

	"github.com/endrio-maciel/rockeat-go.git/internal/jsonutils"
	"github.com/gorilla/csrf"
)

func (api *Api) HandleGetCSRFtoken(w http.ResponseWriter, r *http.Request) {
	token := csrf.Token(r)
	jsonutils.EncodeJson(w, r, http.StatusOK, map[string]any{
		"csfr_token": token,
	})
}

func (api *Api) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !api.Sessions.Exists(r.Context(), "AuthenticatedUserId") {
			jsonutils.EncodeJson(w, r, http.StatusUnauthorized, map[string]any{
				"message": "must be logged in",
			})
			return
		}
		next.ServeHTTP(w, r)
	})
}
