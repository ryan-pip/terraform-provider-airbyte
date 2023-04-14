// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSmartengageSmartengageEnum string

const (
	SourceSmartengageSmartengageEnumSmartengage SourceSmartengageSmartengageEnum = "smartengage"
)

func (e *SourceSmartengageSmartengageEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "smartengage":
		*e = SourceSmartengageSmartengageEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartengageSmartengageEnum: %s", s)
	}
}

// SourceSmartengage - The values required to configure the source.
type SourceSmartengage struct {
	// API Key
	APIKey     string                           `json:"api_key"`
	SourceType SourceSmartengageSmartengageEnum `json:"sourceType"`
}
