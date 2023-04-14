// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceRkiCovidRkiCovidEnum string

const (
	SourceRkiCovidRkiCovidEnumRkiCovid SourceRkiCovidRkiCovidEnum = "rki-covid"
)

func (e *SourceRkiCovidRkiCovidEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "rki-covid":
		*e = SourceRkiCovidRkiCovidEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRkiCovidRkiCovidEnum: %s", s)
	}
}

// SourceRkiCovid - The values required to configure the source.
type SourceRkiCovid struct {
	SourceType SourceRkiCovidRkiCovidEnum `json:"sourceType"`
	// UTC date in the format 2017-01-25. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}
