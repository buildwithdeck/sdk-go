// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package apierrors

import (
	"encoding/json"
	"github.com/buildwithdeck/sdk-go/models/components"
)

type ErrorMessageResponse struct {
	// A broad categorization of the error. Safe for programmatic use.
	ErrorCategory *components.ErrorCategoryEnum `json:"error_category,omitempty"`
	// The particular error code. Safe for programmatic use.
	ErrorCode *components.ErrorCodeEnum `json:"error_code,omitempty"`
	// A developer-friendly representation of the error code. This may change over time and is not safe for programmatic use.
	ErrorMessage *string `json:"error_message,omitempty"`
	// A user-friendly representation of the error code. null if the error is not related to user action. This may change over time and is not safe for programmatic use.
	DisplayMessage *string                 `json:"display_message,omitempty"`
	HTTPMeta       components.HTTPMetadata `json:"-"`
}

var _ error = &ErrorMessageResponse{}

func (e *ErrorMessageResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
