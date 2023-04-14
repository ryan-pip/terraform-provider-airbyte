// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGetlagoGetlagoEnum string

const (
	SourceGetlagoGetlagoEnumGetlago SourceGetlagoGetlagoEnum = "getlago"
)

func (e *SourceGetlagoGetlagoEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "getlago":
		*e = SourceGetlagoGetlagoEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGetlagoGetlagoEnum: %s", s)
	}
}

// SourceGetlago - The values required to configure the source.
type SourceGetlago struct {
	// Your API Key. See <a href="https://doc.getlago.com/docs/api/intro">here</a>.
	APIKey     string                   `json:"api_key"`
	SourceType SourceGetlagoGetlagoEnum `json:"sourceType"`
}
