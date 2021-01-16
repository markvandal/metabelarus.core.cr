package crsign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func handleMsgCreateSignatureList(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSignatureList) (*sdk.Result, error) {
	signature := k.GetSignature(ctx, msg.RootSignatureId)

	if signature.Service != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "One tries to create a signature list for somebody's signature") // If not, throw an error
	}

	list := &types.SignatureList{
		RootSignatureId: msg.RootSignatureId,
		LastSignatureId: msg.RootSignatureId,
		NextSignatureId: msg.RootSignatureId,
		Metadata:        "{}",
	}

	k.CreateSignatureList(ctx, list)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			"mbcorecr.sign:create.list",
			sdk.NewAttribute(
				"list_id",
				list.Id,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgUpdateSignatureList(ctx sdk.Context, k keeper.Keeper, msg *types.MsgUpdateSignatureList) (*sdk.Result, error) {
	var signatureList = types.SignatureList{
		Id:              msg.Id,
		RootSignatureId: msg.RootSignatureId,
		LastSignatureId: msg.LastSignatureId,
		NextSignatureId: msg.NextSignatureId,
		Metadata:        msg.Metadata,
	}

	// if msg.Creator != k.GetSignatureListOwner(ctx, msg.Id) { // Checks if the the msg sender is the same as the current owner
	//     return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect owner") // If not, throw an error
	// }

	k.UpdateSignatureList(ctx, signatureList)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
