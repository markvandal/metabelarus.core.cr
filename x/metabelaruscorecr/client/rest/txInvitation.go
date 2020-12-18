package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/markvandal/metabelaruscorecr/x/metabelaruscorecr/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createInvitationRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	InviterId string `json:"inviterId"`
	IdentityId string `json:"identityId"`
	CreationDate string `json:"creationDate"`
	ActivationPubKey string `json:"activationPubKey"`
	
}

func createInvitationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createInvitationRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedInviterId := req.InviterId
		
		parsedIdentityId := req.IdentityId
		
		parsedCreationDate := req.CreationDate
		
		parsedActivationPubKey := req.ActivationPubKey
		

		msg := types.NewMsgCreateInvitation(
			creator,
			parsedInviterId,
			parsedIdentityId,
			parsedCreationDate,
			parsedActivationPubKey,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setInvitationRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	InviterId string `json:"inviterId"`
	IdentityId string `json:"identityId"`
	CreationDate string `json:"creationDate"`
	ActivationPubKey string `json:"activationPubKey"`
	
}

func setInvitationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setInvitationRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		
		parsedInviterId := req.InviterId
		
		parsedIdentityId := req.IdentityId
		
		parsedCreationDate := req.CreationDate
		
		parsedActivationPubKey := req.ActivationPubKey
		

		msg := types.NewMsgSetInvitation(
			creator,
			req.ID,
			parsedInviterId,
			parsedIdentityId,
			parsedCreationDate,
			parsedActivationPubKey,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteInvitationRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteInvitationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteInvitationRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		msg := types.NewMsgDeleteInvitation(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
