package types

const (
	// ModuleName defines the module name
	ModuleName = "crsign"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

    // QuerierRoute defines the module's query routing key
    QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
    return []byte(p)
}

const (
	SignatureKey= "signature-value-"
	SignatureCountKey= "signature-count-"
)

const (
	Id2SignKey= "id2sign-value-"
	Id2SignCountKey= "id2sign-count-"
)

const (
	AuthKey= "auth-value-"
	AuthCountKey= "auth-count-"
)

const (
	Id2AuthKey= "id2auth-value-"
	Id2AuthCountKey= "id2auth-count-"
)
