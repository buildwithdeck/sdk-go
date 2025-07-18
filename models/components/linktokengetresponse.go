// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// LinkTokenGetResponseLanguage - The language used for the Link session corresponding to the provided Link token
type LinkTokenGetResponseLanguage string

const (
	LinkTokenGetResponseLanguageEn LinkTokenGetResponseLanguage = "EN"
	LinkTokenGetResponseLanguageEs LinkTokenGetResponseLanguage = "ES"
	LinkTokenGetResponseLanguageFr LinkTokenGetResponseLanguage = "FR"
	LinkTokenGetResponseLanguageDe LinkTokenGetResponseLanguage = "DE"
	LinkTokenGetResponseLanguageIt LinkTokenGetResponseLanguage = "IT"
	LinkTokenGetResponseLanguagePt LinkTokenGetResponseLanguage = "PT"
	LinkTokenGetResponseLanguageNl LinkTokenGetResponseLanguage = "NL"
	LinkTokenGetResponseLanguagePl LinkTokenGetResponseLanguage = "PL"
	LinkTokenGetResponseLanguageSv LinkTokenGetResponseLanguage = "SV"
	LinkTokenGetResponseLanguageDa LinkTokenGetResponseLanguage = "DA"
	LinkTokenGetResponseLanguageNo LinkTokenGetResponseLanguage = "NO"
	LinkTokenGetResponseLanguageEt LinkTokenGetResponseLanguage = "ET"
	LinkTokenGetResponseLanguageLt LinkTokenGetResponseLanguage = "LT"
	LinkTokenGetResponseLanguageLv LinkTokenGetResponseLanguage = "LV"
	LinkTokenGetResponseLanguageRo LinkTokenGetResponseLanguage = "RO"
)

func (e LinkTokenGetResponseLanguage) ToPointer() *LinkTokenGetResponseLanguage {
	return &e
}
func (e *LinkTokenGetResponseLanguage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EN":
		fallthrough
	case "ES":
		fallthrough
	case "FR":
		fallthrough
	case "DE":
		fallthrough
	case "IT":
		fallthrough
	case "PT":
		fallthrough
	case "NL":
		fallthrough
	case "PL":
		fallthrough
	case "SV":
		fallthrough
	case "DA":
		fallthrough
	case "NO":
		fallthrough
	case "ET":
		fallthrough
	case "LT":
		fallthrough
	case "LV":
		fallthrough
	case "RO":
		*e = LinkTokenGetResponseLanguage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LinkTokenGetResponseLanguage: %v", v)
	}
}

// LinkTokenGetResponseBetaSourceStatus - The beta source status filter from which data sources will be displayed during the Link session corresponding to the provided Link token.
type LinkTokenGetResponseBetaSourceStatus string

const (
	LinkTokenGetResponseBetaSourceStatusLiveAndBeta LinkTokenGetResponseBetaSourceStatus = "LiveAndBeta"
	LinkTokenGetResponseBetaSourceStatusOnlyLive    LinkTokenGetResponseBetaSourceStatus = "OnlyLive"
	LinkTokenGetResponseBetaSourceStatusOnlyBeta    LinkTokenGetResponseBetaSourceStatus = "OnlyBeta"
)

func (e LinkTokenGetResponseBetaSourceStatus) ToPointer() *LinkTokenGetResponseBetaSourceStatus {
	return &e
}
func (e *LinkTokenGetResponseBetaSourceStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "LiveAndBeta":
		fallthrough
	case "OnlyLive":
		fallthrough
	case "OnlyBeta":
		*e = LinkTokenGetResponseBetaSourceStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LinkTokenGetResponseBetaSourceStatus: %v", v)
	}
}

type LinkTokenGetResponse struct {
	// The language used for the Link session corresponding to the provided Link token
	Language *LinkTokenGetResponseLanguage `json:"language,omitempty"`
	// The beta source status filter from which data sources will be displayed during the Link session corresponding to the provided Link token.
	BetaSourceStatus *LinkTokenGetResponseBetaSourceStatus `json:"beta_source_status,omitempty"`
	// The countries filter from which data sources will be displayed during the Link session corresponding to the provided Link token
	Countries []string `json:"countries,omitempty"`
	// The source types filter from which data sources will be displayed during the Link session corresponding to the provided Link token
	SourceTypes []string `json:"source_types,omitempty"`
	SourceGuids []string `json:"source_guids,omitempty"`
}

func (o *LinkTokenGetResponse) GetLanguage() *LinkTokenGetResponseLanguage {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *LinkTokenGetResponse) GetBetaSourceStatus() *LinkTokenGetResponseBetaSourceStatus {
	if o == nil {
		return nil
	}
	return o.BetaSourceStatus
}

func (o *LinkTokenGetResponse) GetCountries() []string {
	if o == nil {
		return nil
	}
	return o.Countries
}

func (o *LinkTokenGetResponse) GetSourceTypes() []string {
	if o == nil {
		return nil
	}
	return o.SourceTypes
}

func (o *LinkTokenGetResponse) GetSourceGuids() []string {
	if o == nil {
		return nil
	}
	return o.SourceGuids
}
