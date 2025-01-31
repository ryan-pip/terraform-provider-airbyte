// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceLinkedinPagesAuthenticationAccessTokenAuthMethod string

const (
	SourceLinkedinPagesAuthenticationAccessTokenAuthMethodAccessToken SourceLinkedinPagesAuthenticationAccessTokenAuthMethod = "access_token"
)

func (e SourceLinkedinPagesAuthenticationAccessTokenAuthMethod) ToPointer() *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod {
	return &e
}

func (e *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinPagesAuthenticationAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesAuthenticationAccessTokenAuthMethod: %v", v)
	}
}

type SourceLinkedinPagesAuthenticationAccessToken struct {
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	AccessToken string                                                  `json:"access_token"`
	AuthMethod  *SourceLinkedinPagesAuthenticationAccessTokenAuthMethod `json:"auth_method,omitempty"`
}

type SourceLinkedinPagesAuthenticationOAuth20AuthMethod string

const (
	SourceLinkedinPagesAuthenticationOAuth20AuthMethodOAuth20 SourceLinkedinPagesAuthenticationOAuth20AuthMethod = "oAuth2.0"
)

func (e SourceLinkedinPagesAuthenticationOAuth20AuthMethod) ToPointer() *SourceLinkedinPagesAuthenticationOAuth20AuthMethod {
	return &e
}

func (e *SourceLinkedinPagesAuthenticationOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinPagesAuthenticationOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesAuthenticationOAuth20AuthMethod: %v", v)
	}
}

type SourceLinkedinPagesAuthenticationOAuth20 struct {
	AuthMethod *SourceLinkedinPagesAuthenticationOAuth20AuthMethod `json:"auth_method,omitempty"`
	// The client ID of the LinkedIn developer application.
	ClientID string `json:"client_id"`
	// The client secret of the LinkedIn developer application.
	ClientSecret string `json:"client_secret"`
	// The token value generated using the LinkedIn Developers OAuth Token Tools. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-pages/">docs</a> to obtain yours.
	RefreshToken string `json:"refresh_token"`
}

type SourceLinkedinPagesAuthenticationType string

const (
	SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationOAuth20     SourceLinkedinPagesAuthenticationType = "source-linkedin-pages_Authentication_OAuth2.0"
	SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationAccessToken SourceLinkedinPagesAuthenticationType = "source-linkedin-pages_Authentication_Access token"
)

type SourceLinkedinPagesAuthentication struct {
	SourceLinkedinPagesAuthenticationOAuth20     *SourceLinkedinPagesAuthenticationOAuth20
	SourceLinkedinPagesAuthenticationAccessToken *SourceLinkedinPagesAuthenticationAccessToken

	Type SourceLinkedinPagesAuthenticationType
}

func CreateSourceLinkedinPagesAuthenticationSourceLinkedinPagesAuthenticationOAuth20(sourceLinkedinPagesAuthenticationOAuth20 SourceLinkedinPagesAuthenticationOAuth20) SourceLinkedinPagesAuthentication {
	typ := SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationOAuth20

	return SourceLinkedinPagesAuthentication{
		SourceLinkedinPagesAuthenticationOAuth20: &sourceLinkedinPagesAuthenticationOAuth20,
		Type:                                     typ,
	}
}

func CreateSourceLinkedinPagesAuthenticationSourceLinkedinPagesAuthenticationAccessToken(sourceLinkedinPagesAuthenticationAccessToken SourceLinkedinPagesAuthenticationAccessToken) SourceLinkedinPagesAuthentication {
	typ := SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationAccessToken

	return SourceLinkedinPagesAuthentication{
		SourceLinkedinPagesAuthenticationAccessToken: &sourceLinkedinPagesAuthenticationAccessToken,
		Type: typ,
	}
}

func (u *SourceLinkedinPagesAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceLinkedinPagesAuthenticationOAuth20 := new(SourceLinkedinPagesAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinPagesAuthenticationOAuth20); err == nil {
		u.SourceLinkedinPagesAuthenticationOAuth20 = sourceLinkedinPagesAuthenticationOAuth20
		u.Type = SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationOAuth20
		return nil
	}

	sourceLinkedinPagesAuthenticationAccessToken := new(SourceLinkedinPagesAuthenticationAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinPagesAuthenticationAccessToken); err == nil {
		u.SourceLinkedinPagesAuthenticationAccessToken = sourceLinkedinPagesAuthenticationAccessToken
		u.Type = SourceLinkedinPagesAuthenticationTypeSourceLinkedinPagesAuthenticationAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLinkedinPagesAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinPagesAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceLinkedinPagesAuthenticationOAuth20)
	}

	if u.SourceLinkedinPagesAuthenticationAccessToken != nil {
		return json.Marshal(u.SourceLinkedinPagesAuthenticationAccessToken)
	}

	return nil, nil
}

type SourceLinkedinPagesLinkedinPages string

const (
	SourceLinkedinPagesLinkedinPagesLinkedinPages SourceLinkedinPagesLinkedinPages = "linkedin-pages"
)

func (e SourceLinkedinPagesLinkedinPages) ToPointer() *SourceLinkedinPagesLinkedinPages {
	return &e
}

func (e *SourceLinkedinPagesLinkedinPages) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linkedin-pages":
		*e = SourceLinkedinPagesLinkedinPages(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinPagesLinkedinPages: %v", v)
	}
}

type SourceLinkedinPages struct {
	Credentials *SourceLinkedinPagesAuthentication `json:"credentials,omitempty"`
	// Specify the Organization ID
	OrgID      string                           `json:"org_id"`
	SourceType SourceLinkedinPagesLinkedinPages `json:"sourceType"`
}
