// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/buildwithdeck/sdk-go/models/components"
)

type PostLinkConnectResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LinkConnectResponse *components.LinkConnectResponse
}

func (o *PostLinkConnectResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostLinkConnectResponse) GetLinkConnectResponse() *components.LinkConnectResponse {
	if o == nil {
		return nil
	}
	return o.LinkConnectResponse
}
