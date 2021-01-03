package mbcorecr

import (
	"encoding/hex"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	// https://godoc.org/github.com/decred/dcrd/dcrec/secp256k1#example-package--EncryptDecryptMessage

	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func handleMsgCreateSuperIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSuperIdentity) (*sdk.Result, error) {
	inviteHelper, err := NewInviteHelper(msg.Creator, &ctx, &k.BankKeeper, &k.AuthKeeper)
	if err != nil {
		return nil, err
	}

	if err := inviteHelper.WithDenom(types.SuperInviteDenom); err != nil {
		return nil, err
	}

	inviteAcc, err := NewInviteAccount(msg.Uid, ctx, k.AuthKeeper)
	if err != nil {
		return nil, err
	}

	if err := inviteHelper.Pay(); err != nil {
		return nil, err
	}

	inviteAddr, err := sdk.AccAddressFromBech32(inviteAcc.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrNewAccount, err.Error())
	}

	if err := k.BankKeeper.SetBalances(
		ctx,
		inviteAddr,
		types.SuperIdentityCoinsPack,
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrNewAccount, err.Error())
	}

	// Create an identity
	identityKey := k.CreateIdentity(ctx, types.Identity{
		AccountID:    inviteAcc.PubKey,
		IdentityType: types.IdentityType_CITIZEN,
		Details:      "",
		InvitationId: "",
		CreationDt:   msg.CreationDt,
	})

	// Add 1.000.000 simple tokens
	// Add 1.000 stake

	resp := &types.IdentityAccount{
		Uid:        inviteAcc.Uid,
		Address:    inviteAcc.Address,
		Mnemonic:   inviteAcc.Mnemonic,
		PubKey:     inviteAcc.PubKey,
		PrivKey:    inviteAcc.PrivKey,
		IdentityID: identityKey,
	}

	payload, err := json.Marshal(resp)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCipher, err.Error())
	}

	pubKey, err := inviteHelper.GetPubKey()
	if err != nil {
		return nil, err
	}
	cipherPayload, err := mbutils.EncryptPayload(pubKey.Bytes(), payload)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCipher, err.Error())
	}

	ecnryptedPayload := hex.EncodeToString(cipherPayload)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventGovCreateIdentity,
			sdk.NewAttribute(
				types.EventAttrIdentityType,
				string(types.AttrIdentityTypeSuper),
			),
			sdk.NewAttribute(
				types.EventAttrIdentityPayload,
				ecnryptedPayload,
			),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
