package rest

import (
	"net/http"
	"strconv"

    "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
    "github.com/gorilla/mux"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createSignatureRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Identity string `json:"identity"`
	Service string `json:"service"`
	Key string `json:"key"`
	Secret string `json:"secret"`
	CreationDt string `json:"creationDt"`
	AvailabilityDt string `json:"availabilityDt"`
	
}

func createSignatureHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSignatureRequest
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
		
		parsedService := req.Service
		
		parsedKey := req.Key
		
		parsedSecret := req.Secret
		
		parsedCreationDt := req.CreationDt
		
		parsedAvailabilityDt := req.AvailabilityDt
		

		msg := types.NewMsgCreateSignature(
			req.Creator,
			parsedIdentity,
			parsedService,
			parsedKey,
			parsedSecret,
			parsedCreationDt,
			parsedAvailabilityDt,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateSignatureRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	Identity string `json:"identity"`
	Service string `json:"service"`
	Key string `json:"key"`
	Secret string `json:"secret"`
	CreationDt string `json:"creationDt"`
	AvailabilityDt string `json:"availabilityDt"`
	
}


func updateSignatureHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		var req updateSignatureRequest
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
		
		parsedService := req.Service
		
		parsedKey := req.Key
		
		parsedSecret := req.Secret
		
		parsedCreationDt := req.CreationDt
		
		parsedAvailabilityDt := req.AvailabilityDt
		

		msg := types.NewMsgUpdateSignature(
			req.Creator,
            id,
			parsedIdentity,
			parsedService,
			parsedKey,
			parsedSecret,
			parsedCreationDt,
			parsedAvailabilityDt,
			
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteSignatureRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
}

func deleteSignatureHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
        id := mux.Vars(r)["id"]

		var req deleteSignatureRequest
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

		msg := types.NewMsgDeleteSignature(
			req.Creator,
            id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
