package api

import (
	"errors"
	"net/http"

	"github.com/endrio-maciel/rockeat-go.git/internal/jsonutils"
	"github.com/endrio-maciel/rockeat-go.git/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (api *Api) HandleSubscribeUserToAuction(w http.ResponseWriter, r *http.Request) {
	rawProductId := chi.URLParam(r, "product_id")

	productId, err := uuid.Parse(rawProductId)
	if err != nil {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"message": "invalid product id - must be a valid uuid",
		})
		return
	}

	_, err = api.ProductService.GetProductById(r.Context(), productId)
	if err != nil {
		if errors.Is(err, services.ErrProductNotFound) {
			jsonutils.EncodeJson(w, r, http.StatusNotFound, map[string]any{
				"message": "no product with given id",
			})
			return
		}
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "unexpected error, try again later",
		})
		return
	}

	_, ok := api.Sessions.Get(r.Context(), "AuthenticatedUserId").(uuid.UUID) //userId
	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
			"message": "unexpected error, try again later",
		})
		return
	}

	api.AuctionLobby.Lock()
	// _, ok := api.AuctionLobby.Rooms[productId] // room
	api.AuctionLobby.Unlock()

	if !ok {
		jsonutils.EncodeJson(w, r, http.StatusBadRequest, map[string]any{
			"message": "the auction has ended",
		})
		return
	}

	// conn, err = api.WsUpgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	jsonutils.EncodeJson(w, r, http.StatusInternalServerError, map[string]any{
	// 		"message": "could not upgrade connection to a websocket protocol",
	// 	})
	// 	return
	// }

	// client := services.NewClient(room, conn, userId)

	// room.Resgister <- client
	// go client.ReadEventLoop()
	// go client.WriteEventLoop()
	for {

	}

}
