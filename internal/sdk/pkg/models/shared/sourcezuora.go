// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceZuoraZuoraEnum string

const (
	SourceZuoraZuoraEnumZuora SourceZuoraZuoraEnum = "zuora"
)

func (e *SourceZuoraZuoraEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "zuora":
		*e = SourceZuoraZuoraEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZuoraZuoraEnum: %s", s)
	}
}

// SourceZuora - The values required to configure the source.
type SourceZuora struct {
	// Client ID
	ClientID string `json:"client_id"`
	// Client Secret
	ClientSecret string `json:"client_secret"`
	// Defines whether use the SANDBOX or PRODUCTION environment.
	IsSandbox  *bool                `json:"is_sandbox,omitempty"`
	SourceType SourceZuoraZuoraEnum `json:"sourceType"`
	// Start Date in format: YYYY-MM-DD
	StartDate string `json:"start_date"`
	// The amount of days for each data-chunk begining from start_date. Bigger the value - faster the fetch. (Min=1, as for a Day; Max=364, as for a Year).
	WindowInDays *int64 `json:"window_in_days,omitempty"`
}
