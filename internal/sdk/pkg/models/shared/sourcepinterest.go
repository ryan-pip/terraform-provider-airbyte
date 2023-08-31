// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourcePinterestAuthorizationMethodAccessTokenAuthMethod string

const (
	SourcePinterestAuthorizationMethodAccessTokenAuthMethodAccessToken SourcePinterestAuthorizationMethodAccessTokenAuthMethod = "access_token"
)

func (e SourcePinterestAuthorizationMethodAccessTokenAuthMethod) ToPointer() *SourcePinterestAuthorizationMethodAccessTokenAuthMethod {
	return &e
}

func (e *SourcePinterestAuthorizationMethodAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourcePinterestAuthorizationMethodAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestAuthorizationMethodAccessTokenAuthMethod: %v", v)
	}
}

type SourcePinterestAuthorizationMethodAccessToken struct {
	// The Access Token to make authenticated requests.
	AccessToken string                                                  `json:"access_token"`
	AuthMethod  SourcePinterestAuthorizationMethodAccessTokenAuthMethod `json:"auth_method"`
}

type SourcePinterestAuthorizationMethodOAuth20AuthMethod string

const (
	SourcePinterestAuthorizationMethodOAuth20AuthMethodOauth20 SourcePinterestAuthorizationMethodOAuth20AuthMethod = "oauth2.0"
)

func (e SourcePinterestAuthorizationMethodOAuth20AuthMethod) ToPointer() *SourcePinterestAuthorizationMethodOAuth20AuthMethod {
	return &e
}

func (e *SourcePinterestAuthorizationMethodOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourcePinterestAuthorizationMethodOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestAuthorizationMethodOAuth20AuthMethod: %v", v)
	}
}

type SourcePinterestAuthorizationMethodOAuth20 struct {
	AuthMethod SourcePinterestAuthorizationMethodOAuth20AuthMethod `json:"auth_method"`
	// The Client ID of your OAuth application
	ClientID *string `json:"client_id,omitempty"`
	// The Client Secret of your OAuth application.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Refresh Token to obtain new Access Token, when it's expired.
	RefreshToken string `json:"refresh_token"`
}

type SourcePinterestAuthorizationMethodType string

const (
	SourcePinterestAuthorizationMethodTypeSourcePinterestAuthorizationMethodOAuth20     SourcePinterestAuthorizationMethodType = "source-pinterest_Authorization Method_OAuth2.0"
	SourcePinterestAuthorizationMethodTypeSourcePinterestAuthorizationMethodAccessToken SourcePinterestAuthorizationMethodType = "source-pinterest_Authorization Method_Access Token"
)

type SourcePinterestAuthorizationMethod struct {
	SourcePinterestAuthorizationMethodOAuth20     *SourcePinterestAuthorizationMethodOAuth20
	SourcePinterestAuthorizationMethodAccessToken *SourcePinterestAuthorizationMethodAccessToken

	Type SourcePinterestAuthorizationMethodType
}

func CreateSourcePinterestAuthorizationMethodSourcePinterestAuthorizationMethodOAuth20(sourcePinterestAuthorizationMethodOAuth20 SourcePinterestAuthorizationMethodOAuth20) SourcePinterestAuthorizationMethod {
	typ := SourcePinterestAuthorizationMethodTypeSourcePinterestAuthorizationMethodOAuth20

	return SourcePinterestAuthorizationMethod{
		SourcePinterestAuthorizationMethodOAuth20: &sourcePinterestAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourcePinterestAuthorizationMethodSourcePinterestAuthorizationMethodAccessToken(sourcePinterestAuthorizationMethodAccessToken SourcePinterestAuthorizationMethodAccessToken) SourcePinterestAuthorizationMethod {
	typ := SourcePinterestAuthorizationMethodTypeSourcePinterestAuthorizationMethodAccessToken

	return SourcePinterestAuthorizationMethod{
		SourcePinterestAuthorizationMethodAccessToken: &sourcePinterestAuthorizationMethodAccessToken,
		Type: typ,
	}
}

func (u *SourcePinterestAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourcePinterestAuthorizationMethodOAuth20 := new(SourcePinterestAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourcePinterestAuthorizationMethodOAuth20); err == nil {
		u.SourcePinterestAuthorizationMethodOAuth20 = sourcePinterestAuthorizationMethodOAuth20
		u.Type = SourcePinterestAuthorizationMethodTypeSourcePinterestAuthorizationMethodOAuth20
		return nil
	}

	sourcePinterestAuthorizationMethodAccessToken := new(SourcePinterestAuthorizationMethodAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourcePinterestAuthorizationMethodAccessToken); err == nil {
		u.SourcePinterestAuthorizationMethodAccessToken = sourcePinterestAuthorizationMethodAccessToken
		u.Type = SourcePinterestAuthorizationMethodTypeSourcePinterestAuthorizationMethodAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourcePinterestAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourcePinterestAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourcePinterestAuthorizationMethodOAuth20)
	}

	if u.SourcePinterestAuthorizationMethodAccessToken != nil {
		return json.Marshal(u.SourcePinterestAuthorizationMethodAccessToken)
	}

	return nil, nil
}

type SourcePinterestPinterest string

const (
	SourcePinterestPinterestPinterest SourcePinterestPinterest = "pinterest"
)

func (e SourcePinterestPinterest) ToPointer() *SourcePinterestPinterest {
	return &e
}

func (e *SourcePinterestPinterest) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pinterest":
		*e = SourcePinterestPinterest(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestPinterest: %v", v)
	}
}

type SourcePinterestStatus string

const (
	SourcePinterestStatusActive   SourcePinterestStatus = "ACTIVE"
	SourcePinterestStatusPaused   SourcePinterestStatus = "PAUSED"
	SourcePinterestStatusArchived SourcePinterestStatus = "ARCHIVED"
)

func (e SourcePinterestStatus) ToPointer() *SourcePinterestStatus {
	return &e
}

func (e *SourcePinterestStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "PAUSED":
		fallthrough
	case "ARCHIVED":
		*e = SourcePinterestStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePinterestStatus: %v", v)
	}
}

type SourcePinterest struct {
	Credentials *SourcePinterestAuthorizationMethod `json:"credentials,omitempty"`
	SourceType  SourcePinterestPinterest            `json:"sourceType"`
	// A date in the format YYYY-MM-DD. If you have not set a date, it would be defaulted to latest allowed date by api (89 days from today).
	StartDate types.Date `json:"start_date"`
	// Entity statuses based off of campaigns, ad_groups, and ads. If you do not have a status set, it will be ignored completely.
	Status []SourcePinterestStatus `json:"status,omitempty"`
}
