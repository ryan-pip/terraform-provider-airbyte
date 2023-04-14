// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePunkAPIPunkAPIEnum string

const (
	SourcePunkAPIPunkAPIEnumPunkAPI SourcePunkAPIPunkAPIEnum = "punk-api"
)

func (e *SourcePunkAPIPunkAPIEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "punk-api":
		*e = SourcePunkAPIPunkAPIEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePunkAPIPunkAPIEnum: %s", s)
	}
}

// SourcePunkAPI - The values required to configure the source.
type SourcePunkAPI struct {
	// To extract specific data with Unique ID
	BrewedAfter string `json:"brewed_after"`
	// To extract specific data with Unique ID
	BrewedBefore string `json:"brewed_before"`
	// To extract specific data with Unique ID
	ID         *string                  `json:"id,omitempty"`
	SourceType SourcePunkAPIPunkAPIEnum `json:"sourceType"`
}
