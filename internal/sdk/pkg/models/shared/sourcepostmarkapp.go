// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourcePostmarkappPostmarkapp string

const (
	SourcePostmarkappPostmarkappPostmarkapp SourcePostmarkappPostmarkapp = "postmarkapp"
)

func (e SourcePostmarkappPostmarkapp) ToPointer() *SourcePostmarkappPostmarkapp {
	return &e
}

func (e *SourcePostmarkappPostmarkapp) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "postmarkapp":
		*e = SourcePostmarkappPostmarkapp(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePostmarkappPostmarkapp: %v", v)
	}
}

type SourcePostmarkapp struct {
	// API Key for account
	XPostmarkAccountToken string `json:"X-Postmark-Account-Token"`
	// API Key for server
	XPostmarkServerToken string                       `json:"X-Postmark-Server-Token"`
	SourceType           SourcePostmarkappPostmarkapp `json:"sourceType"`
}
