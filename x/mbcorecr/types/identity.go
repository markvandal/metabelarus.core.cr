package types

var _ IdentityI = &Identity{}

type IdentityI interface {
	ExportId() string
	VerifyIdentityType(identityType IdentityType) bool
}

func (this Identity) ExportId() string {
	return this.Id
}

func (this Identity) VerifyIdentityType(identityType IdentityType) bool {
	return this.IdentityType == identityType
}
