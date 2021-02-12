package helper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/keeper"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

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
	k *keeper.Keeper,
) (*InviteHelper, error) {
	inviterAddr, err := sdk.AccAddressFromBech32(InviterBech32)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrCreator, err.Error())
	}

	inviteHelper := &InviteHelper{
		InviterAddress: sdk.AccAddress(inviterAddr),
		bankKeeper:     &k.BankKeeper,
		authKeeper:     &k.AuthKeeper,
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
	if err := (*this.bankKeeper).SubtractCoins(
		*this.ctx,
		this.InviterAddress,
		sdk.Coins{*this.coin},
	); err != nil {
		return sdkerrors.Wrap(types.ErrCreator, err.Error())
	}

	return nil
}
