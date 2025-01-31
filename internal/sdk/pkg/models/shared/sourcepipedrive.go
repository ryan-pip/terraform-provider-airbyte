// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourcePipedriveAPIKeyAuthenticationAuthType string

const (
	SourcePipedriveAPIKeyAuthenticationAuthTypeToken SourcePipedriveAPIKeyAuthenticationAuthType = "Token"
)

func (e SourcePipedriveAPIKeyAuthenticationAuthType) ToPointer() *SourcePipedriveAPIKeyAuthenticationAuthType {
	return &e
}

func (e *SourcePipedriveAPIKeyAuthenticationAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourcePipedriveAPIKeyAuthenticationAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePipedriveAPIKeyAuthenticationAuthType: %v", v)
	}
}

type SourcePipedriveAPIKeyAuthentication struct {
	// The Pipedrive API Token.
	APIToken string                                      `json:"api_token"`
	AuthType SourcePipedriveAPIKeyAuthenticationAuthType `json:"auth_type"`
}

type SourcePipedrivePipedrive string

const (
	SourcePipedrivePipedrivePipedrive SourcePipedrivePipedrive = "pipedrive"
)

func (e SourcePipedrivePipedrive) ToPointer() *SourcePipedrivePipedrive {
	return &e
}

func (e *SourcePipedrivePipedrive) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pipedrive":
		*e = SourcePipedrivePipedrive(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePipedrivePipedrive: %v", v)
	}
}

type SourcePipedrive struct {
	Authorization *SourcePipedriveAPIKeyAuthentication `json:"authorization,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated. When specified and not None, then stream will behave as incremental
	ReplicationStartDate time.Time                `json:"replication_start_date"`
	SourceType           SourcePipedrivePipedrive `json:"sourceType"`
}
