package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
    "github.com/gorilla/mux"
)

func listAuthHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/list-auth", types.QuerierRoute), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func getAuthHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		res, height, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/get-auth/%s", types.QuerierRoute, id), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}
