// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGainsightPxGainsightPx string

const (
	SourceGainsightPxGainsightPxGainsightPx SourceGainsightPxGainsightPx = "gainsight-px"
)

func (e SourceGainsightPxGainsightPx) ToPointer() *SourceGainsightPxGainsightPx {
	return &e
}

func (e *SourceGainsightPxGainsightPx) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gainsight-px":
		*e = SourceGainsightPxGainsightPx(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGainsightPxGainsightPx: %v", v)
	}
}

type SourceGainsightPx struct {
	// The Aptrinsic API Key which is recieved from the dashboard settings (ref - https://app.aptrinsic.com/settings/api-keys)
	APIKey     string                       `json:"api_key"`
	SourceType SourceGainsightPxGainsightPx `json:"sourceType"`
}
