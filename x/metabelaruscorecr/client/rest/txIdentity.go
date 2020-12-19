package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createIdentityRequest struct {
	BaseReq       rest.BaseReq `json:"base_req"`
	AccountID     string       `json:"accountID"`
	Details       string       `json:"details"`
	IdenitityType string       `json:"idenitityType"`
	AuthPubKey    string       `json:"authPubKey"`
}

func createIdentityHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createIdentityRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		parsedAccountID, err := sdk.AccAddressFromBech32(req.AccountID)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedDetails := req.Details

		parsedIdenitityType, err := strconv.Atoi(req.IdenitityType)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())

			return
		}

		parsedAuthPubKey := req.AuthPubKey

		msg := types.NewMsgCreateIdentity(
			parsedAccountID,
			types.IdentityType(parsedIdenitityType),
			parsedDetails,
			parsedAuthPubKey,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setIdentityRequest struct { // @TODO Rewrite structure to actual type
	BaseReq       rest.BaseReq `json:"base_req"`
	ID            string       `json:"id"`
	Details       string       `json:"details"`
	IdenitityType string       `json:"idenitityType"`
	AuthPubKey    string       `json:"authPubKey"`
}

func setIdentityHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setIdentityRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		parsedDetails := req.Details

		parsedIdenitityType, err := strconv.Atoi(req.IdenitityType)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())

			return
		}

		parsedAuthPubKey := req.AuthPubKey

		msg := types.NewMsgSetIdentity(
			req.ID,
			types.IdentityType(parsedIdenitityType),
			parsedDetails,
			parsedAuthPubKey,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
