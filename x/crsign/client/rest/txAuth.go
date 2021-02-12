package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type authRequest struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Creator  string       `json:"creator"`
	Identity string       `json:"identity"`
	Key      string       `json:"key"`
}

func requestAuthHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req authRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedIdentity := req.Identity

		parsedKey := req.Key

		msg := types.NewMsgRequestAuth(
			parsedIdentity,
			req.Creator,
			parsedKey,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type authConfirmation struct {
	BaseReq  rest.BaseReq `json:"base_req"`
	Identity string       `json:"identity"`
	Service  string       `json:"service"`
}

func confirmAuthHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req authConfirmation
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Identity)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedService := req.Service

		msg := types.NewMsgConfirmAuth(
			req.Identity,
			parsedService,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
