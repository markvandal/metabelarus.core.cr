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
	if m.RecordType == RecordType_IDENTITY_MUTUAL_RECORD {
		if m.GetParentId() == "" {
			return m.GetIdentity()
		}
	}

	return ""
}
