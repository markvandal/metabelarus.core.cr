package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createSignatureListRequest struct {
	BaseReq         rest.BaseReq `json:"base_req"`
	Creator         string       `json:"creator"`
	RootSignatureId string       `json:"rootSignatureId"`
}

func createSignatureListHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSignatureListRequest
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

		parsedRootSignatureId := req.RootSignatureId

		msg := types.NewMsgCreateSignatureList(
			req.Creator,
			parsedRootSignatureId,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateSignatureListRequest struct {
	BaseReq         rest.BaseReq `json:"base_req"`
	Creator         string       `json:"creator"`
	RootSignatureId string       `json:"rootSignatureId"`
	LastSignatureId string       `json:"lastSignatureId"`
	NextSignatureId string       `json:"nextSignatureId"`
	Metadata        string       `json:"metadata"`
}

func updateSignatureListHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req updateSignatureListRequest
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

		parsedRootSignatureId := req.RootSignatureId

		parsedLastSignatureId := req.LastSignatureId

		parsedNextSignatureId := req.NextSignatureId

		parsedMetadata := req.Metadata

		msg := types.NewMsgUpdateSignatureList(
			req.Creator,
			id,
			parsedRootSignatureId,
			parsedLastSignatureId,
			parsedNextSignatureId,
			parsedMetadata,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
