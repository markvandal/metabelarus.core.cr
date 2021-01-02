package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type updateIdentityRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	AccountID string       `json:"AccountID"`
	Details   string       `json:"Details"`
}

func updateIdentityHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		var req updateIdentityRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.AccountID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedAccountID := req.AccountID

		parsedDetails := req.Details

		msg := types.NewMsgUpdateIdentity(
			id,
			parsedAccountID,
			parsedDetails,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
