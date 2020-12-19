package types

const (
	// ModuleName is the name of the module
	ModuleName = "metabelaruscorecr"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for routing msgs
	RouterKey = ModuleName

	// QuerierRoute to be used for querier msgs
	QuerierRoute = ModuleName
)

const (
	IdentityPrefix = "identity-"
)

const (
	InvitationPrefix = "invitation-"
)

const (
	ConfirmationPrefix = "confirmation-"
)
