// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType string

const (
	SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthTypeAPIKey SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType = "api_key"
)

func (e SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType) ToPointer() *SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType {
	return &e
}

func (e *SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_key":
		*e = SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType: %v", v)
	}
}

type SourceSalesloftCredentialsAuthenticateViaAPIKey struct {
	// API Key for making authenticated requests. More instruction on how to find this value in our <a href="https://docs.airbyte.com/integrations/sources/salesloft#setup-guide">docs</a>
	APIKey   string                                                  `json:"api_key"`
	AuthType SourceSalesloftCredentialsAuthenticateViaAPIKeyAuthType `json:"auth_type"`
}

type SourceSalesloftCredentialsAuthenticateViaOAuthAuthType string

const (
	SourceSalesloftCredentialsAuthenticateViaOAuthAuthTypeOauth20 SourceSalesloftCredentialsAuthenticateViaOAuthAuthType = "oauth2.0"
)

func (e SourceSalesloftCredentialsAuthenticateViaOAuthAuthType) ToPointer() *SourceSalesloftCredentialsAuthenticateViaOAuthAuthType {
	return &e
}

func (e *SourceSalesloftCredentialsAuthenticateViaOAuthAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceSalesloftCredentialsAuthenticateViaOAuthAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesloftCredentialsAuthenticateViaOAuthAuthType: %v", v)
	}
}

type SourceSalesloftCredentialsAuthenticateViaOAuth struct {
	// Access Token for making authenticated requests.
	AccessToken string                                                 `json:"access_token"`
	AuthType    SourceSalesloftCredentialsAuthenticateViaOAuthAuthType `json:"auth_type"`
	// The Client ID of your Salesloft developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Salesloft developer application.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining a new access token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceSalesloftCredentialsType string

const (
	SourceSalesloftCredentialsTypeSourceSalesloftCredentialsAuthenticateViaOAuth  SourceSalesloftCredentialsType = "source-salesloft_Credentials_Authenticate via OAuth"
	SourceSalesloftCredentialsTypeSourceSalesloftCredentialsAuthenticateViaAPIKey SourceSalesloftCredentialsType = "source-salesloft_Credentials_Authenticate via API Key"
)

type SourceSalesloftCredentials struct {
	SourceSalesloftCredentialsAuthenticateViaOAuth  *SourceSalesloftCredentialsAuthenticateViaOAuth
	SourceSalesloftCredentialsAuthenticateViaAPIKey *SourceSalesloftCredentialsAuthenticateViaAPIKey

	Type SourceSalesloftCredentialsType
}

func CreateSourceSalesloftCredentialsSourceSalesloftCredentialsAuthenticateViaOAuth(sourceSalesloftCredentialsAuthenticateViaOAuth SourceSalesloftCredentialsAuthenticateViaOAuth) SourceSalesloftCredentials {
	typ := SourceSalesloftCredentialsTypeSourceSalesloftCredentialsAuthenticateViaOAuth

	return SourceSalesloftCredentials{
		SourceSalesloftCredentialsAuthenticateViaOAuth: &sourceSalesloftCredentialsAuthenticateViaOAuth,
		Type: typ,
	}
}

func CreateSourceSalesloftCredentialsSourceSalesloftCredentialsAuthenticateViaAPIKey(sourceSalesloftCredentialsAuthenticateViaAPIKey SourceSalesloftCredentialsAuthenticateViaAPIKey) SourceSalesloftCredentials {
	typ := SourceSalesloftCredentialsTypeSourceSalesloftCredentialsAuthenticateViaAPIKey

	return SourceSalesloftCredentials{
		SourceSalesloftCredentialsAuthenticateViaAPIKey: &sourceSalesloftCredentialsAuthenticateViaAPIKey,
		Type: typ,
	}
}

func (u *SourceSalesloftCredentials) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSalesloftCredentialsAuthenticateViaOAuth := new(SourceSalesloftCredentialsAuthenticateViaOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSalesloftCredentialsAuthenticateViaOAuth); err == nil {
		u.SourceSalesloftCredentialsAuthenticateViaOAuth = sourceSalesloftCredentialsAuthenticateViaOAuth
		u.Type = SourceSalesloftCredentialsTypeSourceSalesloftCredentialsAuthenticateViaOAuth
		return nil
	}

	sourceSalesloftCredentialsAuthenticateViaAPIKey := new(SourceSalesloftCredentialsAuthenticateViaAPIKey)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSalesloftCredentialsAuthenticateViaAPIKey); err == nil {
		u.SourceSalesloftCredentialsAuthenticateViaAPIKey = sourceSalesloftCredentialsAuthenticateViaAPIKey
		u.Type = SourceSalesloftCredentialsTypeSourceSalesloftCredentialsAuthenticateViaAPIKey
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSalesloftCredentials) MarshalJSON() ([]byte, error) {
	if u.SourceSalesloftCredentialsAuthenticateViaOAuth != nil {
		return json.Marshal(u.SourceSalesloftCredentialsAuthenticateViaOAuth)
	}

	if u.SourceSalesloftCredentialsAuthenticateViaAPIKey != nil {
		return json.Marshal(u.SourceSalesloftCredentialsAuthenticateViaAPIKey)
	}

	return nil, nil
}

type SourceSalesloftSalesloft string

const (
	SourceSalesloftSalesloftSalesloft SourceSalesloftSalesloft = "salesloft"
)

func (e SourceSalesloftSalesloft) ToPointer() *SourceSalesloftSalesloft {
	return &e
}

func (e *SourceSalesloftSalesloft) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "salesloft":
		*e = SourceSalesloftSalesloft(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSalesloftSalesloft: %v", v)
	}
}

type SourceSalesloft struct {
	Credentials SourceSalesloftCredentials `json:"credentials"`
	SourceType  SourceSalesloftSalesloft   `json:"sourceType"`
	// The date from which you'd like to replicate data for Salesloft API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
}
