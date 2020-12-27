package mbcorecr

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

func handleMsgCreateSuperIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSuperIdentity) (*sdk.Result, error) {
	//k.CreateIdentity(ctx, *msg)

	kring, err := keyring.New(
		sdk.KeyringServiceName(),
		keyring.BackendMemory,
		"",
		nil,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	var (
		info     keyring.Info
		mnemonic string
		pkey     string
		pub      string
	)

	uid := uuid.New().String()

	path := "m/44'/0'/0'/0/0"

	_, mnemonic, err = kring.NewMnemonic(uid, keyring.English, path, hd.Secp256k1)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	info, err = kring.NewAccount(uid, mnemonic, "", path, hd.Secp256k1)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	// @TODO pkey should be reencrypted with a password
	pkey, err = kring.ExportPrivKeyArmor(uid, "")
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	pub, err = sdk.Bech32ifyPubKey(sdk.Bech32PubKeyTypeAccPub, info.GetPubKey())
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	// @TODO Pack the account details in JSON and encrypt it
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventGovCreateIdentity,
			sdk.NewAttribute(types.EventAttrIdentityType, string(types.AttrIdentityTypeSuper)),
			sdk.NewAttribute(types.EventAttrIdentityUid, uid),
			sdk.NewAttribute(types.EventAttrIdentityAddress, info.GetAddress().String()),
			sdk.NewAttribute(types.EventAttrIdentityMnemonic, mnemonic),
			sdk.NewAttribute(types.EventAttrIdentityPubKey, pub),
			sdk.NewAttribute(types.EventAttrIdentityPrivKey, pkey),
		),
	)

	// Create an account

	// Add 1.000.000 simple tokens
	// Add 1.000 stake
	// Add token 100 invitesuper
	// Add token 150 invite0
	// Add token 100 invite1
	// Add token 50 invite2
	// Add token 15 invite3
	// Add token 5 invite4

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
