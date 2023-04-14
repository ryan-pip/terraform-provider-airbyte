// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceHubspotAuthenticationPrivateAppCredentialsEnum - Name of the credentials set
type SourceHubspotAuthenticationPrivateAppCredentialsEnum string

const (
	SourceHubspotAuthenticationPrivateAppCredentialsEnumPrivateAppCredentials SourceHubspotAuthenticationPrivateAppCredentialsEnum = "Private App Credentials"
)

func (e *SourceHubspotAuthenticationPrivateAppCredentialsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Private App Credentials":
		*e = SourceHubspotAuthenticationPrivateAppCredentialsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotAuthenticationPrivateAppCredentialsEnum: %s", s)
	}
}

// SourceHubspotAuthenticationPrivateApp - Choose how to authenticate to HubSpot.
type SourceHubspotAuthenticationPrivateApp struct {
	// HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.
	AccessToken string `json:"access_token"`
	// Name of the credentials set
	CredentialsTitle SourceHubspotAuthenticationPrivateAppCredentialsEnum `json:"credentials_title"`
}

// SourceHubspotAuthenticationOAuthCredentialsEnum - Name of the credentials
type SourceHubspotAuthenticationOAuthCredentialsEnum string

const (
	SourceHubspotAuthenticationOAuthCredentialsEnumOAuthCredentials SourceHubspotAuthenticationOAuthCredentialsEnum = "OAuth Credentials"
)

func (e *SourceHubspotAuthenticationOAuthCredentialsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "OAuth Credentials":
		*e = SourceHubspotAuthenticationOAuthCredentialsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotAuthenticationOAuthCredentialsEnum: %s", s)
	}
}

// SourceHubspotAuthenticationOAuth - Choose how to authenticate to HubSpot.
type SourceHubspotAuthenticationOAuth struct {
	// The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.
	ClientID string `json:"client_id"`
	// The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.
	ClientSecret string `json:"client_secret"`
	// Name of the credentials
	CredentialsTitle SourceHubspotAuthenticationOAuthCredentialsEnum `json:"credentials_title"`
	// Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.
	RefreshToken string `json:"refresh_token"`
}

type SourceHubspotAuthenticationType string

const (
	SourceHubspotAuthenticationTypeSourceHubspotAuthenticationOAuth      SourceHubspotAuthenticationType = "source-hubspot_Authentication_OAuth"
	SourceHubspotAuthenticationTypeSourceHubspotAuthenticationPrivateApp SourceHubspotAuthenticationType = "source-hubspot_Authentication_Private App"
)

type SourceHubspotAuthentication struct {
	SourceHubspotAuthenticationOAuth      *SourceHubspotAuthenticationOAuth
	SourceHubspotAuthenticationPrivateApp *SourceHubspotAuthenticationPrivateApp

	Type SourceHubspotAuthenticationType
}

func CreateSourceHubspotAuthenticationSourceHubspotAuthenticationOAuth(sourceHubspotAuthenticationOAuth SourceHubspotAuthenticationOAuth) SourceHubspotAuthentication {
	typ := SourceHubspotAuthenticationTypeSourceHubspotAuthenticationOAuth

	return SourceHubspotAuthentication{
		SourceHubspotAuthenticationOAuth: &sourceHubspotAuthenticationOAuth,
		Type:                             typ,
	}
}

func CreateSourceHubspotAuthenticationSourceHubspotAuthenticationPrivateApp(sourceHubspotAuthenticationPrivateApp SourceHubspotAuthenticationPrivateApp) SourceHubspotAuthentication {
	typ := SourceHubspotAuthenticationTypeSourceHubspotAuthenticationPrivateApp

	return SourceHubspotAuthentication{
		SourceHubspotAuthenticationPrivateApp: &sourceHubspotAuthenticationPrivateApp,
		Type:                                  typ,
	}
}

func (u *SourceHubspotAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceHubspotAuthenticationOAuth := new(SourceHubspotAuthenticationOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHubspotAuthenticationOAuth); err == nil {
		u.SourceHubspotAuthenticationOAuth = sourceHubspotAuthenticationOAuth
		u.Type = SourceHubspotAuthenticationTypeSourceHubspotAuthenticationOAuth
		return nil
	}

	sourceHubspotAuthenticationPrivateApp := new(SourceHubspotAuthenticationPrivateApp)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHubspotAuthenticationPrivateApp); err == nil {
		u.SourceHubspotAuthenticationPrivateApp = sourceHubspotAuthenticationPrivateApp
		u.Type = SourceHubspotAuthenticationTypeSourceHubspotAuthenticationPrivateApp
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceHubspotAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceHubspotAuthenticationOAuth != nil {
		return json.Marshal(u.SourceHubspotAuthenticationOAuth)
	}

	if u.SourceHubspotAuthenticationPrivateApp != nil {
		return json.Marshal(u.SourceHubspotAuthenticationPrivateApp)
	}

	return nil, nil
}

type SourceHubspotHubspotEnum string

const (
	SourceHubspotHubspotEnumHubspot SourceHubspotHubspotEnum = "hubspot"
)

func (e *SourceHubspotHubspotEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "hubspot":
		*e = SourceHubspotHubspotEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotHubspotEnum: %s", s)
	}
}

// SourceHubspot - The values required to configure the source.
type SourceHubspot struct {
	// Choose how to authenticate to HubSpot.
	Credentials SourceHubspotAuthentication `json:"credentials"`
	SourceType  SourceHubspotHubspotEnum    `json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}
