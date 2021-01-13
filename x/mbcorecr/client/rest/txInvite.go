package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createInviteRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Inviter      string       `json:"inviter"`
	Level        string       `json:"level"`
	IdentityType string       `json:"identity_type"`
}

func createInviteHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createInviteRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Inviter)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		identityLevel, ok := types.IdentityLevel_value[req.Level]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Sprintf("Identity level: %s does not exist", req.Level),
			)
			return
		}

		identityType, ok := types.IdentityType_value[req.IdentityType]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Sprintf("Identity type: %s does not exist", req.IdentityType),
			)
			return
		}

		msg := types.NewMsgCreateInvite(
			req.Inviter,
			types.IdentityLevel(identityLevel),
			types.IdentityType(identityType),
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
