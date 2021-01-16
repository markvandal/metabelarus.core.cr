package helper

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"

	"github.com/cosmos/cosmos-sdk/client"
)

type ClientInviteHelper struct {
	clientCtx *client.Context
	Mnemonics string
	Uid       string
	info      keyring.Info
}

func NewClientInviteHelper(ctx *client.Context) *ClientInviteHelper {
	helper := &ClientInviteHelper{
		clientCtx: ctx,
		Mnemonics: "",
		Uid:       "",
	}

	return helper
}

func (this *ClientInviteHelper) ProduceAccount() error {
	mnemonics, err := this.getMnemonics()
	if err != nil {
		return err
	}

	info, err := this.clientCtx.Keyring.NewAccount(
		this.Uid,
		mnemonics,
		types.UnsecureNewAcctountPKPassword,
		types.DefaultWalletPath,
		hd.Secp256k1,
	)
	if err != nil {
		return err
	}

	this.info = info

	return nil
}

func (this *ClientInviteHelper) GetAddressString() string {
	return this.info.GetAddress().String()
}

func (this *ClientInviteHelper) GetPubKeyString() string {
	return sdk.MustBech32ifyPubKey(
		sdk.Bech32PubKeyTypeAccPub,
		this.info.GetPubKey(),
	)
}

func (this *ClientInviteHelper) getMnemonics() (string, error) {
	if this.Mnemonics == "" {
		this.Uid = uuid.New().String()
		kring, err := getTmpKeyring()
		if err != nil {
			return "", err
		}

		_, mnemonics, err := kring.NewMnemonic(
			uuid.New().String(),
			keyring.English,
			types.DefaultWalletPath,
			hd.Secp256k1,
		)
		if err != nil {
			return "", sdkerrors.Wrap(types.ErrKeyringAccount, err.Error())
		}

		this.Mnemonics = mnemonics
	}

	return this.Mnemonics, nil
}

func getTmpKeyring() (keyring.Keyring, error) {
	kring, err := keyring.New(
		sdk.KeyringServiceName(),
		keyring.BackendMemory,
		"",
		nil,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	return kring, nil
}
