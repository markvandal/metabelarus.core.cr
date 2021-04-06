package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
)

func listRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// rest.WriteErrorResponse(w, http.StatusNotFound, "lololo g")
		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/list-request", types.QuerierRoute), nil)
		if err != nil {
			// rest.WriteErrorResponse(w, http.StatusNotFound, "ololo")
			rest.WriteErrorResponse(w, http.StatusNotFound, types.QuerierRoute+"__>"+err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func getRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/get-request/%s", types.QuerierRoute, id), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
