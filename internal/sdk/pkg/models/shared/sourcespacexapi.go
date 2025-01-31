// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceSpacexAPISpacexAPI string

const (
	SourceSpacexAPISpacexAPISpacexAPI SourceSpacexAPISpacexAPI = "spacex-api"
)

func (e SourceSpacexAPISpacexAPI) ToPointer() *SourceSpacexAPISpacexAPI {
	return &e
}

func (e *SourceSpacexAPISpacexAPI) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "spacex-api":
		*e = SourceSpacexAPISpacexAPI(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSpacexAPISpacexAPI: %v", v)
	}
}

type SourceSpacexAPI struct {
	ID         *string                   `json:"id,omitempty"`
	Options    *string                   `json:"options,omitempty"`
	SourceType *SourceSpacexAPISpacexAPI `json:"sourceType,omitempty"`
}
