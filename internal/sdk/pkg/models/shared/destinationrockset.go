// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationRocksetRocksetEnum string

const (
	DestinationRocksetRocksetEnumRockset DestinationRocksetRocksetEnum = "rockset"
)

func (e *DestinationRocksetRocksetEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "rockset":
		*e = DestinationRocksetRocksetEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationRocksetRocksetEnum: %s", s)
	}
}

// DestinationRockset - The values required to configure the destination.
type DestinationRockset struct {
	// Rockset api key
	APIKey string `json:"api_key"`
	// Rockset api URL
	APIServer       *string                       `json:"api_server,omitempty"`
	DestinationType DestinationRocksetRocksetEnum `json:"destinationType"`
	// The Rockset workspace in which collections will be created + written to.
	Workspace string `json:"workspace"`
}
