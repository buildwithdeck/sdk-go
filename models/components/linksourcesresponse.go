// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LinkSourcesResponse struct {
	// List of data sources found.
	Sources []LinkSourcesResponseSource `json:"sources,omitempty"`
}

func (o *LinkSourcesResponse) GetSources() []LinkSourcesResponseSource {
	if o == nil {
		return nil
	}
	return o.Sources
}
