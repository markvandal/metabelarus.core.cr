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
	RecordKey      = "record-value-"
	RecordCountKey = "record-count-"

	Id2RecordKey    = "id2record-value-"
	Id2KeyRecordKey = "id2keyrecord-value-"
)

const (
	AuthKey      = "auth-value-"
	AuthCountKey = "auth-count-"

	Id2ServicesKey = "id2service-value-"
)
