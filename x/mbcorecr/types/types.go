package types

const (
	EventGovCreateIdentity = "mbcorecr.gov:create.identity"
)

const (
	EventAttrIdentityType     = "type"
	EventAttrIdentityUid      = "uid"
	EventAttrIdentityAddress  = "address"
	EventAttrIdentityPubKey   = "pubkey"
	EventAttrIdentityPrivKey  = "privkey"
	EventAttrIdentityMnemonic = "mnemonic"
)

type AttrIdentityType string

const (
	AttrIdentityTypeSuper AttrIdentityType = "super"
)
