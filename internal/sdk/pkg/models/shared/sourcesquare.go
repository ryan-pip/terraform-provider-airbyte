// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceSquareAuthenticationAPIKeyAuthType string

const (
	SourceSquareAuthenticationAPIKeyAuthTypeAPIKey SourceSquareAuthenticationAPIKeyAuthType = "API Key"
)

func (e SourceSquareAuthenticationAPIKeyAuthType) ToPointer() *SourceSquareAuthenticationAPIKeyAuthType {
	return &e
}

func (e *SourceSquareAuthenticationAPIKeyAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "API Key":
		*e = SourceSquareAuthenticationAPIKeyAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSquareAuthenticationAPIKeyAuthType: %v", v)
	}
}

// SourceSquareAuthenticationAPIKey - Choose how to authenticate to Square.
type SourceSquareAuthenticationAPIKey struct {
	// The API key for a Square application
	APIKey   string                                   `json:"api_key"`
	AuthType SourceSquareAuthenticationAPIKeyAuthType `json:"auth_type"`
}

type SourceSquareAuthenticationOauthAuthenticationAuthType string

const (
	SourceSquareAuthenticationOauthAuthenticationAuthTypeOAuth SourceSquareAuthenticationOauthAuthenticationAuthType = "OAuth"
)

func (e SourceSquareAuthenticationOauthAuthenticationAuthType) ToPointer() *SourceSquareAuthenticationOauthAuthenticationAuthType {
	return &e
}

func (e *SourceSquareAuthenticationOauthAuthenticationAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth":
		*e = SourceSquareAuthenticationOauthAuthenticationAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSquareAuthenticationOauthAuthenticationAuthType: %v", v)
	}
}

// SourceSquareAuthenticationOauthAuthentication - Choose how to authenticate to Square.
type SourceSquareAuthenticationOauthAuthentication struct {
	AuthType SourceSquareAuthenticationOauthAuthenticationAuthType `json:"auth_type"`
	// The Square-issued ID of your application
	ClientID string `json:"client_id"`
	// The Square-issued application secret for your application
	ClientSecret string `json:"client_secret"`
	// A refresh token generated using the above client ID and secret
	RefreshToken string `json:"refresh_token"`
}

type SourceSquareAuthenticationType string

const (
	SourceSquareAuthenticationTypeSourceSquareAuthenticationOauthAuthentication SourceSquareAuthenticationType = "source-square_Authentication_Oauth authentication"
	SourceSquareAuthenticationTypeSourceSquareAuthenticationAPIKey              SourceSquareAuthenticationType = "source-square_Authentication_API key"
)

type SourceSquareAuthentication struct {
	SourceSquareAuthenticationOauthAuthentication *SourceSquareAuthenticationOauthAuthentication
	SourceSquareAuthenticationAPIKey              *SourceSquareAuthenticationAPIKey

	Type SourceSquareAuthenticationType
}

func CreateSourceSquareAuthenticationSourceSquareAuthenticationOauthAuthentication(sourceSquareAuthenticationOauthAuthentication SourceSquareAuthenticationOauthAuthentication) SourceSquareAuthentication {
	typ := SourceSquareAuthenticationTypeSourceSquareAuthenticationOauthAuthentication

	return SourceSquareAuthentication{
		SourceSquareAuthenticationOauthAuthentication: &sourceSquareAuthenticationOauthAuthentication,
		Type: typ,
	}
}

func CreateSourceSquareAuthenticationSourceSquareAuthenticationAPIKey(sourceSquareAuthenticationAPIKey SourceSquareAuthenticationAPIKey) SourceSquareAuthentication {
	typ := SourceSquareAuthenticationTypeSourceSquareAuthenticationAPIKey

	return SourceSquareAuthentication{
		SourceSquareAuthenticationAPIKey: &sourceSquareAuthenticationAPIKey,
		Type:                             typ,
	}
}

func (u *SourceSquareAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSquareAuthenticationOauthAuthentication := new(SourceSquareAuthenticationOauthAuthentication)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSquareAuthenticationOauthAuthentication); err == nil {
		u.SourceSquareAuthenticationOauthAuthentication = sourceSquareAuthenticationOauthAuthentication
		u.Type = SourceSquareAuthenticationTypeSourceSquareAuthenticationOauthAuthentication
		return nil
	}

	sourceSquareAuthenticationAPIKey := new(SourceSquareAuthenticationAPIKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSquareAuthenticationAPIKey); err == nil {
		u.SourceSquareAuthenticationAPIKey = sourceSquareAuthenticationAPIKey
		u.Type = SourceSquareAuthenticationTypeSourceSquareAuthenticationAPIKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSquareAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceSquareAuthenticationOauthAuthentication != nil {
		return json.Marshal(u.SourceSquareAuthenticationOauthAuthentication)
	}

	if u.SourceSquareAuthenticationAPIKey != nil {
		return json.Marshal(u.SourceSquareAuthenticationAPIKey)
	}

	return nil, nil
}

type SourceSquareSquare string

const (
	SourceSquareSquareSquare SourceSquareSquare = "square"
)

func (e SourceSquareSquare) ToPointer() *SourceSquareSquare {
	return &e
}

func (e *SourceSquareSquare) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "square":
		*e = SourceSquareSquare(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSquareSquare: %v", v)
	}
}

type SourceSquare struct {
	// Choose how to authenticate to Square.
	Credentials *SourceSquareAuthentication `json:"credentials,omitempty"`
	// In some streams there is an option to include deleted objects (Items, Categories, Discounts, Taxes)
	IncludeDeletedObjects *bool `json:"include_deleted_objects,omitempty"`
	// Determines whether to use the sandbox or production environment.
	IsSandbox  bool               `json:"is_sandbox"`
	SourceType SourceSquareSquare `json:"sourceType"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. If not set, all data will be replicated.
	StartDate *types.Date `json:"start_date,omitempty"`
}
