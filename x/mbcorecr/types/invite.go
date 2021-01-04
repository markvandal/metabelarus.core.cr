package types

import (
	"encoding/base64"
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	mbutils "github.com/metabelarus/mbcorecr/mb/utils"
)

func (this *InviteAccount) GetAccAddress() (sdk.AccAddress, error) {
	inviteAddr, err := sdk.AccAddressFromBech32(this.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrNewAccount, err.Error())
	}

	return inviteAddr, nil
}

func (this *InviteAccount) SetBalances(ctx sdk.Context, bank BankKeeper, balances sdk.Coins) error {
	inviteAddr, err := this.GetAccAddress()
	if err != nil {
		return err
	}

	if err := bank.SetBalances(
		ctx,
		inviteAddr,
		balances,
	); err != nil {
		return sdkerrors.Wrap(ErrNewAccount, err.Error())
	}

	return nil
}

func (this *InviteAccount) EncryptStr(data string) (string, error) {
	pubKey := sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, this.PubKey)
	ecnrypted, err := mbutils.EncryptPayload(pubKey.Bytes(), []byte(data))
	if err != nil {
		return "", sdkerrors.Wrap(ErrCryptDetails, err.Error())
	}

	return base64.URLEncoding.EncodeToString(ecnrypted), nil
}

func (this *InviteAccount) EncryptData(data interface{}) (string, error) {
	if data == nil {
		data = this
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return "", sdkerrors.Wrap(ErrCipher, err.Error())
	}

	return this.EncryptStr(string(payload))
}
