// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType string

const (
	SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthTypeToken SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType = "Token"
)

func (e SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType) ToPointer() *SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType {
	return &e
}

func (e *SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Token":
		*e = SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType: %v", v)
	}
}

// SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken - Choose how to authenticate to Harvest.
type SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken struct {
	// Log into Harvest and then create new <a href="https://id.getharvest.com/developers"> personal access token</a>.
	APIToken string                                                                                 `json:"api_token"`
	AuthType *SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessTokenAuthType `json:"auth_type,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken

func (c *SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken) UnmarshalJSON(bs []byte) error {
	data := _SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "api_token")
	delete(additionalFields, "auth_type")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType string

const (
	SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthTypeClient SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType = "Client"
)

func (e SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType) ToPointer() *SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType {
	return &e
}

func (e *SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Client":
		*e = SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType: %v", v)
	}
}

// SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth - Choose how to authenticate to Harvest.
type SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth struct {
	AuthType *SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuthAuthType `json:"auth_type,omitempty"`
	// The Client ID of your Harvest developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Harvest developer application.
	ClientSecret string `json:"client_secret"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth

func (c *SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth) UnmarshalJSON(bs []byte) error {
	data := _SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "auth_type")
	delete(additionalFields, "client_id")
	delete(additionalFields, "client_secret")
	delete(additionalFields, "refresh_token")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth(c))
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	bs, err = json.Marshal(c.AdditionalProperties)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(bs), &out); err != nil {
		return nil, err
	}

	return json.Marshal(out)
}

type SourceHarvestUpdateAuthenticationMechanismType string

const (
	SourceHarvestUpdateAuthenticationMechanismTypeSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth         SourceHarvestUpdateAuthenticationMechanismType = "source-harvest-update_Authentication mechanism_Authenticate via Harvest (OAuth)"
	SourceHarvestUpdateAuthenticationMechanismTypeSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceHarvestUpdateAuthenticationMechanismType = "source-harvest-update_Authentication mechanism_Authenticate with Personal Access Token"
)

type SourceHarvestUpdateAuthenticationMechanism struct {
	SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth         *SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth
	SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken *SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken

	Type SourceHarvestUpdateAuthenticationMechanismType
}

func CreateSourceHarvestUpdateAuthenticationMechanismSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth(sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth) SourceHarvestUpdateAuthenticationMechanism {
	typ := SourceHarvestUpdateAuthenticationMechanismTypeSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth

	return SourceHarvestUpdateAuthenticationMechanism{
		SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth: &sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth,
		Type: typ,
	}
}

func CreateSourceHarvestUpdateAuthenticationMechanismSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken(sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken) SourceHarvestUpdateAuthenticationMechanism {
	typ := SourceHarvestUpdateAuthenticationMechanismTypeSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken

	return SourceHarvestUpdateAuthenticationMechanism{
		SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken: &sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceHarvestUpdateAuthenticationMechanism) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth := new(SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth); err == nil {
		u.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth = sourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth
		u.Type = SourceHarvestUpdateAuthenticationMechanismTypeSourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth
		return nil
	}

	sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken := new(SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken); err == nil {
		u.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken = sourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken
		u.Type = SourceHarvestUpdateAuthenticationMechanismTypeSourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceHarvestUpdateAuthenticationMechanism) MarshalJSON() ([]byte, error) {
	if u.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth != nil {
		return json.Marshal(u.SourceHarvestUpdateAuthenticationMechanismAuthenticateViaHarvestOAuth)
	}

	if u.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken != nil {
		return json.Marshal(u.SourceHarvestUpdateAuthenticationMechanismAuthenticateWithPersonalAccessToken)
	}

	return nil, nil
}

type SourceHarvestUpdate struct {
	// Harvest account ID. Required for all Harvest requests in pair with Personal Access Token
	AccountID string `json:"account_id"`
	// Choose how to authenticate to Harvest.
	Credentials *SourceHarvestUpdateAuthenticationMechanism `json:"credentials,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data after this date will not be replicated.
	ReplicationEndDate *time.Time `json:"replication_end_date,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	ReplicationStartDate time.Time `json:"replication_start_date"`
}
