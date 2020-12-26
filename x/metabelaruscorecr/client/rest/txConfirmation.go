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

type createConfirmationRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	IdenitityID string `json:"idenitityID"`
	CreationDate string `json:"creationDate"`
	ExpirationDate string `json:"expirationDate"`
	ConfirmatorID string `json:"confirmatorID"`
	CenterGeo string `json:"centerGeo"`
	Status string `json:"status"`
	NextTryDate string `json:"nextTryDate"`
	
}

func createConfirmationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createConfirmationRequest
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

		
		parsedIdenitityID := req.IdenitityID
		
		parsedCreationDate := req.CreationDate
		
		parsedExpirationDate := req.ExpirationDate
		
		parsedConfirmatorID := req.ConfirmatorID
		
		parsedCenterGeo := req.CenterGeo
		
		parsedStatus := req.Status
		
		parsedNextTryDate := req.NextTryDate
		

		msg := types.NewMsgCreateConfirmation(
			creator,
			parsedIdenitityID,
			parsedCreationDate,
			parsedExpirationDate,
			parsedConfirmatorID,
			parsedCenterGeo,
			parsedStatus,
			parsedNextTryDate,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setConfirmationRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	IdenitityID string `json:"idenitityID"`
	CreationDate string `json:"creationDate"`
	ExpirationDate string `json:"expirationDate"`
	ConfirmatorID string `json:"confirmatorID"`
	CenterGeo string `json:"centerGeo"`
	Status string `json:"status"`
	NextTryDate string `json:"nextTryDate"`
	
}

func setConfirmationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setConfirmationRequest
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

		
		parsedIdenitityID := req.IdenitityID
		
		parsedCreationDate := req.CreationDate
		
		parsedExpirationDate := req.ExpirationDate
		
		parsedConfirmatorID := req.ConfirmatorID
		
		parsedCenterGeo := req.CenterGeo
		
		parsedStatus := req.Status
		
		parsedNextTryDate := req.NextTryDate
		

		msg := types.NewMsgSetConfirmation(
			creator,
			req.ID,
			parsedIdenitityID,
			parsedCreationDate,
			parsedExpirationDate,
			parsedConfirmatorID,
			parsedCenterGeo,
			parsedStatus,
			parsedNextTryDate,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteConfirmationRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteConfirmationHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteConfirmationRequest
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
		msg := types.NewMsgDeleteConfirmation(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
