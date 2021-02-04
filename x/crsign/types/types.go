package types

const (
	EventRequestAuth = "mbcorecr.crsign:request.auth"
	EventConfirmAuth = "mbcorecr.crsign:confirm.auth"

	EventCreateRecord = "mbcorecr.crsign:create.record"
)

const (
	EventAttrAuthId   = "auth_id"
	EventAttrRecordId = "record_id"
)

const (
	DefaultAuthLifeTime = "12h"
)

func IsProviderRecord(recordType RecordType) bool {
	switch recordType {
	case RecordType_PROVIDER_MUTUAL_RECORD:
	case RecordType_PROVIDER_RECORD:
	case RecordType_PROVIDER_SIGNABLE_RECORD:
		return true
	}

	return false
}

func IsIdentityRecord(recordType RecordType) bool {
	return !IsProviderRecord(recordType)
}

func IsMutualRecord(recordType RecordType) bool {
	switch recordType {
	case RecordType_PROVIDER_MUTUAL_RECORD:
	case RecordType_IDENTITY_MUTUAL_RECORD:
		return true
	}

	return false
}
