package api

import (
	"net/http"

	"github.com/endrio-maciel/rockeat-go.git/internal/jsonutils"
	"github.com/endrio-maciel/rockeat-go.git/internal/usecase/user"
)

func (api *Api) HandleSignupUser(w http.ResponseWriter, r *http.Request) {
	data, problems, err := jsonutils.DecodeValidJson[user.CreateUserReq](r)
	if err != nil {
		_ = jsonutils.EncodeJson(w, r, http.StatusUnprocessableEntity, problems)
	}

}

func (api *Api) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO - NOT IMPLEMENTED")
}

func (api *Api) HandleLogoutUser(w http.ResponseWriter, r *http.Request) {
	panic("TODO - NOT IMPLEMENTED")
}
