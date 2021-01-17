
package types


func (this *Auth) GetId() string {
	return this.Identity + ":" + this.Service
}