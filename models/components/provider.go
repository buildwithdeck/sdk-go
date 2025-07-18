// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Provider struct {
	// Unique identifier of the data provider associated to this connectiond
	ID *string `json:"id,omitempty"`
	// Name of the data source provider to this connection
	Name *string `json:"name,omitempty"`
}

func (o *Provider) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Provider) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
