package types

type IdentityI interface {
	ExportId() string
	ExportAddress() string
}

func (this Identity) ExportId() string {
	return this.Id
}

func (this Identity) ExportAddress() string {
	return this.Address
}
