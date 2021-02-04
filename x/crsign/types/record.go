package types

import (
	"strings"
)

func (m Record) GetParentId() string {
	parts := strings.Split(m.Id, ".")
	if len(parts) == 1 {
		return ""
	}

	return parts[0]
}

func (m Record) GetChildId() string {
	if m.RecordType == RecordType_PROVIDER_MUTUAL_RECORD {
		return m.GetId() + "."
	}

	return ""
}

func (m Record) IsChildRecord() bool {
	parts := strings.Split(m.Id, ".")
	if len(parts) == 1 {
		return false
	}

	return true
}

func (m Record) IsParentRecord() bool {
	return IsMutualRecord(m.RecordType)
}
