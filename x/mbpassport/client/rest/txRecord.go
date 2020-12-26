package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/markvandal/metabelaruscorecr/x/mbpassport/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createRecordRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	IdentityId string `json:"IdentityId"`
	ServiceId string `json:"ServiceId"`
	ServiceType string `json:"ServiceType"`
	Key string `json:"Key"`
	UserValue string `json:"UserValue"`
	ServiceValue string `json:"ServiceValue"`
	CreationDt string `json:"CreationDt"`
	UpdateDt string `json:"UpdateDt"`
	
}

func createRecordHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRecordRequest
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

		
		parsedIdentityId := req.IdentityId
		
		parsedServiceId := req.ServiceId
		
		parsedServiceType := req.ServiceType
		
		parsedKey := req.Key
		
		parsedUserValue := req.UserValue
		
		parsedServiceValue := req.ServiceValue
		
		parsedCreationDt := req.CreationDt
		
		parsedUpdateDt := req.UpdateDt
		

		msg := types.NewMsgCreateRecord(
			creator,
			parsedIdentityId,
			parsedServiceId,
			parsedServiceType,
			parsedKey,
			parsedUserValue,
			parsedServiceValue,
			parsedCreationDt,
			parsedUpdateDt,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setRecordRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	ID 		string `json:"id"`
	Creator string `json:"creator"`
	IdentityId string `json:"IdentityId"`
	ServiceId string `json:"ServiceId"`
	ServiceType string `json:"ServiceType"`
	Key string `json:"Key"`
	UserValue string `json:"UserValue"`
	ServiceValue string `json:"ServiceValue"`
	CreationDt string `json:"CreationDt"`
	UpdateDt string `json:"UpdateDt"`
	
}

func setRecordHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setRecordRequest
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

		
		parsedIdentityId := req.IdentityId
		
		parsedServiceId := req.ServiceId
		
		parsedServiceType := req.ServiceType
		
		parsedKey := req.Key
		
		parsedUserValue := req.UserValue
		
		parsedServiceValue := req.ServiceValue
		
		parsedCreationDt := req.CreationDt
		
		parsedUpdateDt := req.UpdateDt
		

		msg := types.NewMsgSetRecord(
			creator,
			req.ID,
			parsedIdentityId,
			parsedServiceId,
			parsedServiceType,
			parsedKey,
			parsedUserValue,
			parsedServiceValue,
			parsedCreationDt,
			parsedUpdateDt,
			
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type deleteRecordRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string `json:"creator"`
	ID 		string `json:"id"`
}

func deleteRecordHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req deleteRecordRequest
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
		msg := types.NewMsgDeleteRecord(req.ID, creator)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
