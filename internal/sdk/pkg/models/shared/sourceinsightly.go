// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceInsightlyInsightlyEnum string

const (
	SourceInsightlyInsightlyEnumInsightly SourceInsightlyInsightlyEnum = "insightly"
)

func (e *SourceInsightlyInsightlyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "insightly":
		*e = SourceInsightlyInsightlyEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceInsightlyInsightlyEnum: %s", s)
	}
}

// SourceInsightly - The values required to configure the source.
type SourceInsightly struct {
	SourceType SourceInsightlyInsightlyEnum `json:"sourceType"`
	// The date from which you'd like to replicate data for Insightly in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated. Note that it will be used only for incremental streams.
	StartDate string `json:"start_date"`
	// Your Insightly API token.
	Token string `json:"token"`
}
