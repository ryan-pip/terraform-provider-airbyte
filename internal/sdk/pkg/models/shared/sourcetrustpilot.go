// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnum string

const (
	SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnumApikey SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnum = "apikey"
)

func (e *SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "apikey":
		*e = SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnum: %s", s)
	}
}

// SourceTrustpilotAuthorizationMethodAPIKey - The API key authentication method gives you access to only the streams which are part of the Public API. When you want to get streams available via the Consumer API (e.g. the private reviews) you need to use authentication method OAuth 2.0.
type SourceTrustpilotAuthorizationMethodAPIKey struct {
	AuthType *SourceTrustpilotAuthorizationMethodAPIKeyAuthTypeEnum `json:"auth_type,omitempty"`
	// The API key of the Trustpilot API application.
	ClientID string `json:"client_id"`
}

type SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnum string

const (
	SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnumOauth20 SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnum = "oauth2.0"
)

func (e *SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "oauth2.0":
		*e = SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnum: %s", s)
	}
}

type SourceTrustpilotAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                                                  `json:"access_token"`
	AuthType    *SourceTrustpilotAuthorizationMethodOAuth20AuthTypeEnum `json:"auth_type,omitempty"`
	// The API key of the Trustpilot API application. (represents the OAuth Client ID)
	ClientID string `json:"client_id"`
	// The Secret of the Trustpilot API application. (represents the OAuth Client Secret)
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access_token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceTrustpilotAuthorizationMethodType string

const (
	SourceTrustpilotAuthorizationMethodTypeSourceTrustpilotAuthorizationMethodOAuth20 SourceTrustpilotAuthorizationMethodType = "source-trustpilot_Authorization Method_OAuth 2.0"
	SourceTrustpilotAuthorizationMethodTypeSourceTrustpilotAuthorizationMethodAPIKey  SourceTrustpilotAuthorizationMethodType = "source-trustpilot_Authorization Method_API Key"
)

type SourceTrustpilotAuthorizationMethod struct {
	SourceTrustpilotAuthorizationMethodOAuth20 *SourceTrustpilotAuthorizationMethodOAuth20
	SourceTrustpilotAuthorizationMethodAPIKey  *SourceTrustpilotAuthorizationMethodAPIKey

	Type SourceTrustpilotAuthorizationMethodType
}

func CreateSourceTrustpilotAuthorizationMethodSourceTrustpilotAuthorizationMethodOAuth20(sourceTrustpilotAuthorizationMethodOAuth20 SourceTrustpilotAuthorizationMethodOAuth20) SourceTrustpilotAuthorizationMethod {
	typ := SourceTrustpilotAuthorizationMethodTypeSourceTrustpilotAuthorizationMethodOAuth20

	return SourceTrustpilotAuthorizationMethod{
		SourceTrustpilotAuthorizationMethodOAuth20: &sourceTrustpilotAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceTrustpilotAuthorizationMethodSourceTrustpilotAuthorizationMethodAPIKey(sourceTrustpilotAuthorizationMethodAPIKey SourceTrustpilotAuthorizationMethodAPIKey) SourceTrustpilotAuthorizationMethod {
	typ := SourceTrustpilotAuthorizationMethodTypeSourceTrustpilotAuthorizationMethodAPIKey

	return SourceTrustpilotAuthorizationMethod{
		SourceTrustpilotAuthorizationMethodAPIKey: &sourceTrustpilotAuthorizationMethodAPIKey,
		Type: typ,
	}
}

func (u *SourceTrustpilotAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceTrustpilotAuthorizationMethodOAuth20 := new(SourceTrustpilotAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceTrustpilotAuthorizationMethodOAuth20); err == nil {
		u.SourceTrustpilotAuthorizationMethodOAuth20 = sourceTrustpilotAuthorizationMethodOAuth20
		u.Type = SourceTrustpilotAuthorizationMethodTypeSourceTrustpilotAuthorizationMethodOAuth20
		return nil
	}

	sourceTrustpilotAuthorizationMethodAPIKey := new(SourceTrustpilotAuthorizationMethodAPIKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceTrustpilotAuthorizationMethodAPIKey); err == nil {
		u.SourceTrustpilotAuthorizationMethodAPIKey = sourceTrustpilotAuthorizationMethodAPIKey
		u.Type = SourceTrustpilotAuthorizationMethodTypeSourceTrustpilotAuthorizationMethodAPIKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceTrustpilotAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceTrustpilotAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceTrustpilotAuthorizationMethodOAuth20)
	}

	if u.SourceTrustpilotAuthorizationMethodAPIKey != nil {
		return json.Marshal(u.SourceTrustpilotAuthorizationMethodAPIKey)
	}

	return nil, nil
}

type SourceTrustpilotTrustpilotEnum string

const (
	SourceTrustpilotTrustpilotEnumTrustpilot SourceTrustpilotTrustpilotEnum = "trustpilot"
)

func (e *SourceTrustpilotTrustpilotEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "trustpilot":
		*e = SourceTrustpilotTrustpilotEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTrustpilotTrustpilotEnum: %s", s)
	}
}

// SourceTrustpilot - The values required to configure the source.
type SourceTrustpilot struct {
	// The names of business units which shall be synchronized. Some streams e.g. configured_business_units or private_reviews use this configuration.
	BusinessUnits []string                            `json:"business_units"`
	Credentials   SourceTrustpilotAuthorizationMethod `json:"credentials"`
	SourceType    SourceTrustpilotTrustpilotEnum      `json:"sourceType"`
	// For streams with sync. method incremental the start date time to be used
	StartDate string `json:"start_date"`
}
