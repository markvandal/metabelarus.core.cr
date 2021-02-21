package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/metabelarus/mbcorecr/x/crconsent/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createRequestRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	Initiator   string       `json:"initiator"`
	Recipient   string       `json:"recipient"`
	RequestType string       `json:"requestType"`
	Status      string       `json:"status"`
	Value       string       `json:"value"`
	Memo        string       `json:"memo"`
	PromoUrl    string       `json:"promoUrl"`
	CreationDt  string       `json:"creationDt"`
	FinalDt     string       `json:"finalDt"`
}

func createRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRequestRequest
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

		parsedInitiator := req.Initiator

		parsedRecipient := req.Recipient

		parsedRequestType := req.RequestType

		parsedStatus := req.Status

		parsedValue64, err := strconv.ParseInt(req.Value, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedValue := int32(parsedValue64)

		parsedMemo := req.Memo

		parsedPromoUrl := req.PromoUrl

		parsedCreationDt := req.CreationDt

		parsedFinalDt := req.FinalDt

		msg := types.NewMsgCreateRequest(
			req.Creator,
			parsedInitiator,
			parsedRecipient,
			parsedRequestType,
			parsedStatus,
			parsedValue,
			parsedMemo,
			parsedPromoUrl,
			parsedCreationDt,
			parsedFinalDt,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateRequestRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	Initiator   string       `json:"initiator"`
	Recipient   string       `json:"recipient"`
	RequestType string       `json:"requestType"`
	Status      string       `json:"status"`
	Value       string       `json:"value"`
	Memo        string       `json:"memo"`
	PromoUrl    string       `json:"promoUrl"`
	CreationDt  string       `json:"creationDt"`
	FinalDt     string       `json:"finalDt"`
}

func updateRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req updateRequestRequest
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

		parsedInitiator := req.Initiator

		parsedRecipient := req.Recipient

		parsedRequestType := req.RequestType

		parsedStatus := req.Status

		parsedValue64, err := strconv.ParseInt(req.Value, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		parsedValue := int32(parsedValue64)

		parsedMemo := req.Memo

		parsedPromoUrl := req.PromoUrl

		parsedCreationDt := req.CreationDt

		parsedFinalDt := req.FinalDt

		msg := types.NewMsgUpdateRequest(
			req.Creator,
			id,
			parsedInitiator,
			parsedRecipient,
			parsedRequestType,
			parsedStatus,
			parsedValue,
			parsedMemo,
			parsedPromoUrl,
			parsedCreationDt,
			parsedFinalDt,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteRequestRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req deleteRequestRequest
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

		msg := types.NewMsgDeleteRequest(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
