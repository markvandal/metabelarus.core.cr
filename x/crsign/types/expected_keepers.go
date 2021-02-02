package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"

	coretype "github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

type IdentityKeeper interface {
	ExportIdentity(ctx sdk.Context, key string) coretype.IdentityI
	HasIdentity(ctx sdk.Context, id string) bool
	GetIdFromAddress(ctx sdk.Context, address string) string
	EnsureIdFromAddress(ctx sdk.Context, address string, creationDt *time.Time) (string, error)
	GetAddressFromId(ctx sdk.Context, id string) string
}

// AccountKeeper is the interface contract that x/auth's keeper implements.
type AccountKeeper interface {
	// Return a new account with the next account number and the specified address. Does not save the new account to the store.
	NewAccountWithAddress(sdk.Context, sdk.AccAddress) authtypes.AccountI

	// Return a new account with the next account number. Does not save the new account to the store.
	NewAccount(sdk.Context, authtypes.AccountI) authtypes.AccountI

	// Retrieve an account from the store.
	GetAccount(sdk.Context, sdk.AccAddress) authtypes.AccountI

	// Set an account in the store.
	SetAccount(sdk.Context, authtypes.AccountI)

	// Remove an account from the store.
	RemoveAccount(sdk.Context, authtypes.AccountI)

	// Iterate over all accounts, calling the provided function. Stop iteraiton when it returns false.
	IterateAccounts(sdk.Context, func(authtypes.AccountI) bool)

	// Fetch the public key of an account at a specified address
	GetPubKey(sdk.Context, sdk.AccAddress) (crypto.PubKey, error)

	// Fetch the sequence of an account at a specified address.
	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)

	// Fetch the next account number, and increment the internal counter.
	GetNextAccountNumber(sdk.Context) uint64
}

// BankViewKeeper defines a module interface that facilitates read only access to
// account balances.
type BankViewKeeper interface {
	ValidateBalance(ctx sdk.Context, addr sdk.AccAddress) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool

	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx sdk.Context) []types.Balance
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	LockedCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	IterateAccountBalances(ctx sdk.Context, addr sdk.AccAddress, cb func(coin sdk.Coin) (stop bool))
	IterateAllBalances(ctx sdk.Context, cb func(address sdk.AccAddress, coin sdk.Coin) (stop bool))
}
