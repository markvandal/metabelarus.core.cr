package helper

import (
	"encoding/base64"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
	"github.com/tendermint/tendermint/crypto"

	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func DeleteAccount(ctx sdk.Context, auth types.AccountKeeper, address string) error {
	tmpAddr, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return sdkerrors.Wrap(types.ErrInvite, err.Error())
	}
	auth.RemoveAccount(ctx, auth.GetAccount(ctx, tmpAddr))

	return nil
}

// NewInviteAccount - Generate payload object for Invite and
// Identity Account responses.
func NewInviteAccount(staticUID string, ctx sdk.Context, authKeeper types.AccountKeeper) (*types.InviteAccount, error) {
	kring, err := keyring.New(
		sdk.KeyringServiceName(),
		keyring.BackendMemory,
		"",
		nil,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyring, err.Error())
	}

	_, mnemonic, err := kring.NewMnemonic(
		uuid.New().String(),
		keyring.English,
		types.DefaultWalletPath,
		hd.Secp256k1,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyringAccount, err.Error())
	}

	info, err := kring.NewAccount(
		staticUID,
		mnemonic,
		types.UnsecureNewAcctountPKPassword,
		types.DefaultWalletPath,
		hd.Secp256k1,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrKeyringAccount, err.Error())
	}

	newAcc := authKeeper.NewAccountWithAddress(ctx, info.GetAddress())
	if err := newAcc.SetPubKey(info.GetPubKey()); err != nil {
		return nil, sdkerrors.Wrap(types.ErrNewAccount, err.Error())
	}
	authKeeper.SetAccount(ctx, newAcc)

	pub, err := sdk.Bech32ifyPubKey(
		sdk.Bech32PubKeyTypeAccPub,
		info.GetPubKey(),
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCryptConversion, err.Error())
	}

	pkey, err := kring.ExportPrivKeyArmor(staticUID, types.UnsecureNewAcctountPKPassword)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCryptConversion, err.Error())
	}

	inviteAcc := &types.InviteAccount{
		Uid:      staticUID,
		Mnemonic: mnemonic,
		Address:  info.GetAddress().String(),
		PubKey:   pub,
		PrivKey:  pkey,
	}

	return inviteAcc, nil
}

type InviteHelper struct {
	Denom          string
	bankKeeper     *types.BankKeeper
	authKeeper     *types.AccountKeeper
	InviterAddress sdk.AccAddress
	coin           *sdk.Coin
	ctx            *sdk.Context
}

func NewInviteHelper(
	InviterBech32 string,
	ctx *sdk.Context,
	bank *types.BankKeeper,
	auth *types.AccountKeeper,
) (*InviteHelper, error) {
	inviterAddr, err := sdk.AccAddressFromBech32(InviterBech32)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreator, err.Error())
	}

	inviteHelper := &InviteHelper{
		InviterAddress: sdk.AccAddress(inviterAddr),
		bankKeeper:     bank,
		authKeeper:     auth,
		ctx:            ctx,
	}

	return inviteHelper, nil
}

func (this *InviteHelper) WithLevel(level types.IdentityLevel) error {
	return this.WithDenom(types.IdentityLevelToDenom[level])
}

func (this *InviteHelper) WithDenom(denom string) error {
	this.coin = &sdk.Coin{
		Denom:  denom,
		Amount: sdk.NewInt(1),
	}

	if !(*this.bankKeeper).HasBalance(*this.ctx, this.InviterAddress, *this.coin) {
		return sdkerrors.Wrap(types.ErrCreator, "there is no super invite")
	}

	return nil
}

func (this *InviteHelper) Pay() error {
	if err := (*this.bankKeeper).SetBalances(
		*this.ctx,
		this.InviterAddress,
		types.SuperIdentityCoinsPack,
	); err != nil {
		return sdkerrors.Wrap(types.ErrNewAccount, err.Error())
	}

	return nil
}

func (this *InviteHelper) GetPubKey() (crypto.PubKey, error) {
	pubKey, err := (*this.authKeeper).GetPubKey(*this.ctx, this.InviterAddress)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreator, err.Error())
	}

	return pubKey, nil
}

func (this *InviteHelper) EncryptStr(data string) (string, error) {
	pubKey, err := this.GetPubKey()
	if err != nil {
		return "", err
	}

	ecnrypted, err := mbutils.EncryptPayload(pubKey.Bytes(), []byte(data))
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrCryptDetails, err.Error())
	}

	return base64.URLEncoding.EncodeToString(ecnrypted), nil
}

func (this *InviteHelper) EncryptData(data interface{}) (string, error) {
	payload, err := json.Marshal(data)
	if err != nil {
		return "", sdkerrors.Wrap(types.ErrCipher, err.Error())
	}

	return this.EncryptStr(string(payload))
}
