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
	SignatureKey= "Signature-value-"
	SignatureCountKey= "Signature-count-"
)

const (
	Id2SignKey= "Id2Sign-value-"
	Id2SignCountKey= "Id2Sign-count-"
)

const (
	AuthKey= "Auth-value-"
	AuthCountKey= "Auth-count-"
)

const (
	Id2AuthKey= "Id2Auth-value-"
	Id2AuthCountKey= "Id2Auth-count-"
)
