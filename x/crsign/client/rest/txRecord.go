package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createRecordRequest struct {
	BaseReq    rest.BaseReq `json:"base_req"`
	Creator    string       `json:"creator"`
	Provider   string       `json:"provider"`
	Key        string       `json:"key"`
	Data       string       `json:"data"`
	Signature  string       `json:"signature"`
	RecordType string       `json:"type"`
	Publicity  string       `json:"publicity"`
	LiveTime   string       `json:"livetime"`
	ParentId   string       `json:"parent"`
}

func createRecordHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRecordRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		liveTime, err := strconv.ParseInt(req.LiveTime, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		publicity, ok := types.PublicityType_value[req.Publicity]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Errorf("Publicity type: %s does not exist", req.Publicity).Error(),
			)
			return
		}

		recordType, ok := types.RecordType_value[req.RecordType]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Errorf("Record type: %s does not exist", req.RecordType).Error(),
			)

			return
		}

		msg := types.NewMsgCreateRecord(
			req.Creator,
			req.Key,
			req.Data,
			req.Signature,
			types.RecordType(recordType),
			types.PublicityType(publicity),
			int32(liveTime),
			req.Provider,
			req.ParentId,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateRecordRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Updater   string       `json:"updater"`
	Data      string       `json:"data"`
	Signature string       `json:"signature"`
	LiveTime  string       `json:"livetime"`
	Action    string       `json:"action"`
}

func updateRecordHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req updateRecordRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		liveTime, err := strconv.ParseInt(req.LiveTime, 10, 32)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		action, ok := types.PublicityType_value[req.Action]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Errorf("Action type: %s does not exist", req.Action).Error(),
			)
			return
		}

		msg := types.NewMsgUpdateRecord(
			req.Updater,
			id,
			req.Data,
			req.Signature,
			int32(liveTime),
			types.RecordUpdate(action),
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteRecordRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Deleter string       `json:"deleter"`
}

func deleteRecordHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req deleteRecordRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgDeleteRecord(
			req.Deleter,
			id,
		)

		err := msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
