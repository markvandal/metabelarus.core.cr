package types

const (
	EventRequestAuth = "mbcorecr.crsign:request.auth"
	EventConfirmAuth = "mbcorecr.crsign:confirm.auth"

	EventCreateRecord = "mbcorecr.crsign:create.record"
	EventUpdateRecord = "mbcorecr.crsign:update.record"
	EventDeleteRecord = "mbcorecr.crsign:delete.record"
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
	case RecordType_PROVIDER_RECORD:
		return true
	case RecordType_PROVIDER_SIGNABLE_RECORD:
		return true
	}

	return false
}

func IsIdentityRecord(recordType RecordType) bool {
	return !IsProviderRecord(recordType)
}
