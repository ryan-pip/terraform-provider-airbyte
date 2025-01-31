// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceZendeskSupportAuthenticationAPITokenCredentials string

const (
	SourceZendeskSupportAuthenticationAPITokenCredentialsAPIToken SourceZendeskSupportAuthenticationAPITokenCredentials = "api_token"
)

func (e SourceZendeskSupportAuthenticationAPITokenCredentials) ToPointer() *SourceZendeskSupportAuthenticationAPITokenCredentials {
	return &e
}

func (e *SourceZendeskSupportAuthenticationAPITokenCredentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "api_token":
		*e = SourceZendeskSupportAuthenticationAPITokenCredentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportAuthenticationAPITokenCredentials: %v", v)
	}
}

// SourceZendeskSupportAuthenticationAPIToken - Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
type SourceZendeskSupportAuthenticationAPIToken struct {
	// The value of the API token generated. See our <a href="https://docs.airbyte.com/integrations/sources/zendesk-support#setup-guide">full documentation</a> for more information on generating this token.
	APIToken    string                                                 `json:"api_token"`
	Credentials *SourceZendeskSupportAuthenticationAPITokenCredentials `json:"credentials,omitempty"`
	// The user email for your Zendesk account.
	Email string `json:"email"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceZendeskSupportAuthenticationAPIToken SourceZendeskSupportAuthenticationAPIToken

func (c *SourceZendeskSupportAuthenticationAPIToken) UnmarshalJSON(bs []byte) error {
	data := _SourceZendeskSupportAuthenticationAPIToken{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceZendeskSupportAuthenticationAPIToken(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "api_token")
	delete(additionalFields, "credentials")
	delete(additionalFields, "email")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceZendeskSupportAuthenticationAPIToken) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceZendeskSupportAuthenticationAPIToken(c))
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

type SourceZendeskSupportAuthenticationOAuth20Credentials string

const (
	SourceZendeskSupportAuthenticationOAuth20CredentialsOauth20 SourceZendeskSupportAuthenticationOAuth20Credentials = "oauth2.0"
)

func (e SourceZendeskSupportAuthenticationOAuth20Credentials) ToPointer() *SourceZendeskSupportAuthenticationOAuth20Credentials {
	return &e
}

func (e *SourceZendeskSupportAuthenticationOAuth20Credentials) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceZendeskSupportAuthenticationOAuth20Credentials(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportAuthenticationOAuth20Credentials: %v", v)
	}
}

// SourceZendeskSupportAuthenticationOAuth20 - Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
type SourceZendeskSupportAuthenticationOAuth20 struct {
	// The OAuth access token. See the <a href="https://developer.zendesk.com/documentation/ticketing/working-with-oauth/creating-and-using-oauth-tokens-with-the-api/">Zendesk docs</a> for more information on generating this token.
	AccessToken string `json:"access_token"`
	// The OAuth client's ID. See <a href="https://docs.searchunify.com/Content/Content-Sources/Zendesk-Authentication-OAuth-Client-ID-Secret.htm#:~:text=Get%20Client%20ID%20and%20Client%20Secret&text=Go%20to%20OAuth%20Clients%20and,will%20be%20displayed%20only%20once.">this guide</a> for more information.
	ClientID *string `json:"client_id,omitempty"`
	// The OAuth client secret. See <a href="https://docs.searchunify.com/Content/Content-Sources/Zendesk-Authentication-OAuth-Client-ID-Secret.htm#:~:text=Get%20Client%20ID%20and%20Client%20Secret&text=Go%20to%20OAuth%20Clients%20and,will%20be%20displayed%20only%20once.">this guide</a> for more information.
	ClientSecret *string                                               `json:"client_secret,omitempty"`
	Credentials  *SourceZendeskSupportAuthenticationOAuth20Credentials `json:"credentials,omitempty"`

	AdditionalProperties interface{} `json:"-"`
}
type _SourceZendeskSupportAuthenticationOAuth20 SourceZendeskSupportAuthenticationOAuth20

func (c *SourceZendeskSupportAuthenticationOAuth20) UnmarshalJSON(bs []byte) error {
	data := _SourceZendeskSupportAuthenticationOAuth20{}

	if err := json.Unmarshal(bs, &data); err != nil {
		return err
	}
	*c = SourceZendeskSupportAuthenticationOAuth20(data)

	additionalFields := make(map[string]interface{})

	if err := json.Unmarshal(bs, &additionalFields); err != nil {
		return err
	}
	delete(additionalFields, "access_token")
	delete(additionalFields, "client_id")
	delete(additionalFields, "client_secret")
	delete(additionalFields, "credentials")

	c.AdditionalProperties = additionalFields

	return nil
}

func (c SourceZendeskSupportAuthenticationOAuth20) MarshalJSON() ([]byte, error) {
	out := map[string]interface{}{}
	bs, err := json.Marshal(_SourceZendeskSupportAuthenticationOAuth20(c))
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

type SourceZendeskSupportAuthenticationType string

const (
	SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAuthenticationOAuth20  SourceZendeskSupportAuthenticationType = "source-zendesk-support_Authentication_OAuth2.0"
	SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAuthenticationAPIToken SourceZendeskSupportAuthenticationType = "source-zendesk-support_Authentication_API Token"
)

type SourceZendeskSupportAuthentication struct {
	SourceZendeskSupportAuthenticationOAuth20  *SourceZendeskSupportAuthenticationOAuth20
	SourceZendeskSupportAuthenticationAPIToken *SourceZendeskSupportAuthenticationAPIToken

	Type SourceZendeskSupportAuthenticationType
}

func CreateSourceZendeskSupportAuthenticationSourceZendeskSupportAuthenticationOAuth20(sourceZendeskSupportAuthenticationOAuth20 SourceZendeskSupportAuthenticationOAuth20) SourceZendeskSupportAuthentication {
	typ := SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAuthenticationOAuth20

	return SourceZendeskSupportAuthentication{
		SourceZendeskSupportAuthenticationOAuth20: &sourceZendeskSupportAuthenticationOAuth20,
		Type: typ,
	}
}

func CreateSourceZendeskSupportAuthenticationSourceZendeskSupportAuthenticationAPIToken(sourceZendeskSupportAuthenticationAPIToken SourceZendeskSupportAuthenticationAPIToken) SourceZendeskSupportAuthentication {
	typ := SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAuthenticationAPIToken

	return SourceZendeskSupportAuthentication{
		SourceZendeskSupportAuthenticationAPIToken: &sourceZendeskSupportAuthenticationAPIToken,
		Type: typ,
	}
}

func (u *SourceZendeskSupportAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceZendeskSupportAuthenticationOAuth20 := new(SourceZendeskSupportAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskSupportAuthenticationOAuth20); err == nil {
		u.SourceZendeskSupportAuthenticationOAuth20 = sourceZendeskSupportAuthenticationOAuth20
		u.Type = SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAuthenticationOAuth20
		return nil
	}

	sourceZendeskSupportAuthenticationAPIToken := new(SourceZendeskSupportAuthenticationAPIToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceZendeskSupportAuthenticationAPIToken); err == nil {
		u.SourceZendeskSupportAuthenticationAPIToken = sourceZendeskSupportAuthenticationAPIToken
		u.Type = SourceZendeskSupportAuthenticationTypeSourceZendeskSupportAuthenticationAPIToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceZendeskSupportAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceZendeskSupportAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceZendeskSupportAuthenticationOAuth20)
	}

	if u.SourceZendeskSupportAuthenticationAPIToken != nil {
		return json.Marshal(u.SourceZendeskSupportAuthenticationAPIToken)
	}

	return nil, nil
}

type SourceZendeskSupportZendeskSupport string

const (
	SourceZendeskSupportZendeskSupportZendeskSupport SourceZendeskSupportZendeskSupport = "zendesk-support"
)

func (e SourceZendeskSupportZendeskSupport) ToPointer() *SourceZendeskSupportZendeskSupport {
	return &e
}

func (e *SourceZendeskSupportZendeskSupport) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "zendesk-support":
		*e = SourceZendeskSupportZendeskSupport(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceZendeskSupportZendeskSupport: %v", v)
	}
}

type SourceZendeskSupport struct {
	// Zendesk allows two authentication methods. We recommend using `OAuth2.0` for Airbyte Cloud users and `API token` for Airbyte Open Source users.
	Credentials *SourceZendeskSupportAuthentication `json:"credentials,omitempty"`
	// Makes each stream read a single page of data.
	IgnorePagination *bool                              `json:"ignore_pagination,omitempty"`
	SourceType       SourceZendeskSupportZendeskSupport `json:"sourceType"`
	// The UTC date and time from which you'd like to replicate data, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// This is your unique Zendesk subdomain that can be found in your account URL. For example, in https://MY_SUBDOMAIN.zendesk.com/, MY_SUBDOMAIN is the value of your subdomain.
	Subdomain string `json:"subdomain"`
}
