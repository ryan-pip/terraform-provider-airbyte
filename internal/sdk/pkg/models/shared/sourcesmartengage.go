// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSmartengageSmartengage string

const (
	SourceSmartengageSmartengageSmartengage SourceSmartengageSmartengage = "smartengage"
)

func (e SourceSmartengageSmartengage) ToPointer() *SourceSmartengageSmartengage {
	return &e
}

func (e *SourceSmartengageSmartengage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "smartengage":
		*e = SourceSmartengageSmartengage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartengageSmartengage: %v", v)
	}
}

type SourceSmartengage struct {
	// API Key
	APIKey     string                       `json:"api_key"`
	SourceType SourceSmartengageSmartengage `json:"sourceType"`
}
