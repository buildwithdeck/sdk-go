// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type RefreshConnectionRequest struct {
	// The access token associated with the Item data is being requested for
	AccessToken    string   `json:"access_token"`
	AccountNumbers []string `json:"account_numbers,omitempty"`
}

func (o *RefreshConnectionRequest) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}

func (o *RefreshConnectionRequest) GetAccountNumbers() []string {
	if o == nil {
		return nil
	}
	return o.AccountNumbers
}
