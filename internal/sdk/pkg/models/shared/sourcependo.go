// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePendoPendo string

const (
	SourcePendoPendoPendo SourcePendoPendo = "pendo"
)

func (e SourcePendoPendo) ToPointer() *SourcePendoPendo {
	return &e
}

func (e *SourcePendoPendo) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pendo":
		*e = SourcePendoPendo(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePendoPendo: %v", v)
	}
}

type SourcePendo struct {
	APIKey     string           `json:"api_key"`
	SourceType SourcePendoPendo `json:"sourceType"`
}
