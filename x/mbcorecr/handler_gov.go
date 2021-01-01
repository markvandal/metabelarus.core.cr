package mbcorecr

import (
	"encoding/hex"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	// https://godoc.org/github.com/decred/dcrd/dcrec/secp256k1#example-package--EncryptDecryptMessage

	"github.com/google/uuid"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func handleMsgCreateSuperIdentity(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreateSuperIdentity) (*sdk.Result, error) {
	kring, err := keyring.New(
		sdk.KeyringServiceName(),
		keyring.BackendMemory,
		"",
		nil,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	pubKey, err := k.AuthKeeper.GetPubKey(ctx, sdk.AccAddress(msg.Creator))
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreator, err.Error())
	}

	coin := &sdk.Coin{
		Denom:  types.SuperInviteDenom,
		Amount: sdk.NewInt(1),
	}

	if !k.BankKeeper.HasBalance(ctx, sdk.AccAddress(msg.Creator), *coin) {
		return nil, sdkerrors.Wrap(types.ErrCreator, "there is no super invite")
	}

	path := types.DefaultWalletPath
	if msg.WalletPath != "" {
		path = msg.WalletPath
	}

	_, mnemonic, err := kring.NewMnemonic(
		uuid.New().String(),
		keyring.English,
		path,
		hd.Secp256k1,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyringAccount, err.Error())
	}

	// ctx.Logger().Info(mnemonic)
	uid := uuid.New().String()
	info, err := kring.NewAccount(
		uid,
		mnemonic,
		types.UnsecureNewAcctountPKPassword,
		path,
		hd.Secp256k1,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyringAccount, err.Error())
	}

	pub, err := sdk.Bech32ifyPubKey(
		sdk.Bech32PubKeyTypeAccPub,
		info.GetPubKey(),
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCryptConversion, err.Error())
	}

	// armoredPub := crypto.ArmorPubKeyBytes(info.GetPubKey().Bytes(), info.GetPubKey().Type())

	// ctx.Logger().Info(pub)
	// ctx.Logger().Info(armoredPub)

	pkey, err := kring.ExportPrivKeyArmor(uid, types.UnsecureNewAcctountPKPassword)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCryptConversion, err.Error())
	}

	// ctx.Logger().Info(pkey)
	// pkUnarm, _, err = crypto.UnarmorDecryptPrivKey(pkey, msg.Password)
	// if err !+ nil {
	// 	return nil, sdkerrors.Wrap(types.ErrCipher, err.Error())
	// }
	// ctx.Logger().Info(string(pkUnarm.Bytes()))

	newAcc := k.AuthKeeper.NewAccountWithAddress(ctx, info.GetAddress())
	if err := newAcc.SetPubKey(info.GetPubKey()); err != nil {
		return nil, sdkerrors.Wrap(types.ErrNewAccount, err.Error())
	}
	k.AuthKeeper.SetAccount(ctx, newAcc)

	if err := k.BankKeeper.SubtractCoins(
		ctx,
		sdk.AccAddress(msg.Creator),
		sdk.Coins{*coin},
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreator, err.Error())
	}

	if err := k.BankKeeper.SetBalances(
		ctx,
		info.GetAddress(),
		types.SuperIdentityCoinsPack,
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrNewAccount, err.Error())
	}

	resp := &types.IdentityAccount{
		Uid:      uid,
		Address:  info.GetAddress().String(),
		Mnemonic: mnemonic,
		PubKey:   pub,
		PrivKey:  pkey,
	}

	payload, err := json.Marshal(resp)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCipher, err.Error())
	}

	// ctx.Logger().Info(string(payload))

	// pubKey, err := sdk.GetPubKeyFromBech32(
	// 	sdk.Bech32PubKeyTypeAccPub,
	// 	msg.PubKey,
	// )
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(types.ErrCryptConversion, err.Error())
	// }

	// pubCipherKeyBytes, err := sdk.GetPubKeyFromBech32(
	// 	sdk.Bech32PubKeyTypeAccPub,
	// 	msg.PubKey,
	// )
	// if err != nil {
	// 	return nil, sdkerrors.Wrap(types.ErrCryptConversion, err.Error())
	// }

	cipherPayload, err := mbutils.EncryptPayload(pubKey.Bytes(), payload)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCipher, err.Error())
	}

	ecnryptedPayload := hex.EncodeToString(cipherPayload)

	// ctx.Logger().Info(string(ecnryptedPayload))

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
