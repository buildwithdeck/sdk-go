// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type AccessTokenRequest struct {
	// The access token associated with the Item data is being requested for
	AccessToken string `json:"access_token"`
}

func (o *AccessTokenRequest) GetAccessToken() string {
	if o == nil {
		return ""
	}
	return o.AccessToken
}
