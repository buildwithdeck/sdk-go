// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/buildwithdeck/sdk-go/models/components"
)

type PostLinkConnectionStatusResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	LinkConnectionStatusResponse *components.LinkConnectionStatusResponse
}

func (o *PostLinkConnectionStatusResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PostLinkConnectionStatusResponse) GetLinkConnectionStatusResponse() *components.LinkConnectionStatusResponse {
	if o == nil {
		return nil
	}
	return o.LinkConnectionStatusResponse
}
