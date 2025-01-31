// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// SourceHubspotUpdateAuthenticationPrivateAppAuthType - Name of the credentials set
type SourceHubspotUpdateAuthenticationPrivateAppAuthType string

const (
	SourceHubspotUpdateAuthenticationPrivateAppAuthTypePrivateAppCredentials SourceHubspotUpdateAuthenticationPrivateAppAuthType = "Private App Credentials"
)

func (e SourceHubspotUpdateAuthenticationPrivateAppAuthType) ToPointer() *SourceHubspotUpdateAuthenticationPrivateAppAuthType {
	return &e
}

func (e *SourceHubspotUpdateAuthenticationPrivateAppAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Private App Credentials":
		*e = SourceHubspotUpdateAuthenticationPrivateAppAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotUpdateAuthenticationPrivateAppAuthType: %v", v)
	}
}

// SourceHubspotUpdateAuthenticationPrivateApp - Choose how to authenticate to HubSpot.
type SourceHubspotUpdateAuthenticationPrivateApp struct {
	// HubSpot Access token. See the <a href="https://developers.hubspot.com/docs/api/private-apps">Hubspot docs</a> if you need help finding this token.
	AccessToken string `json:"access_token"`
	// Name of the credentials set
	CredentialsTitle SourceHubspotUpdateAuthenticationPrivateAppAuthType `json:"credentials_title"`
}

// SourceHubspotUpdateAuthenticationOAuthAuthType - Name of the credentials
type SourceHubspotUpdateAuthenticationOAuthAuthType string

const (
	SourceHubspotUpdateAuthenticationOAuthAuthTypeOAuthCredentials SourceHubspotUpdateAuthenticationOAuthAuthType = "OAuth Credentials"
)

func (e SourceHubspotUpdateAuthenticationOAuthAuthType) ToPointer() *SourceHubspotUpdateAuthenticationOAuthAuthType {
	return &e
}

func (e *SourceHubspotUpdateAuthenticationOAuthAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth Credentials":
		*e = SourceHubspotUpdateAuthenticationOAuthAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceHubspotUpdateAuthenticationOAuthAuthType: %v", v)
	}
}

// SourceHubspotUpdateAuthenticationOAuth - Choose how to authenticate to HubSpot.
type SourceHubspotUpdateAuthenticationOAuth struct {
	// The Client ID of your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this ID.
	ClientID string `json:"client_id"`
	// The client secret for your HubSpot developer application. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this secret.
	ClientSecret string `json:"client_secret"`
	// Name of the credentials
	CredentialsTitle SourceHubspotUpdateAuthenticationOAuthAuthType `json:"credentials_title"`
	// Refresh token to renew an expired access token. See the <a href="https://legacydocs.hubspot.com/docs/methods/oauth2/oauth2-quickstart">Hubspot docs</a> if you need help finding this token.
	RefreshToken string `json:"refresh_token"`
}

type SourceHubspotUpdateAuthenticationType string

const (
	SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateAuthenticationOAuth      SourceHubspotUpdateAuthenticationType = "source-hubspot-update_Authentication_OAuth"
	SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateAuthenticationPrivateApp SourceHubspotUpdateAuthenticationType = "source-hubspot-update_Authentication_Private App"
)

type SourceHubspotUpdateAuthentication struct {
	SourceHubspotUpdateAuthenticationOAuth      *SourceHubspotUpdateAuthenticationOAuth
	SourceHubspotUpdateAuthenticationPrivateApp *SourceHubspotUpdateAuthenticationPrivateApp

	Type SourceHubspotUpdateAuthenticationType
}

func CreateSourceHubspotUpdateAuthenticationSourceHubspotUpdateAuthenticationOAuth(sourceHubspotUpdateAuthenticationOAuth SourceHubspotUpdateAuthenticationOAuth) SourceHubspotUpdateAuthentication {
	typ := SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateAuthenticationOAuth

	return SourceHubspotUpdateAuthentication{
		SourceHubspotUpdateAuthenticationOAuth: &sourceHubspotUpdateAuthenticationOAuth,
		Type:                                   typ,
	}
}

func CreateSourceHubspotUpdateAuthenticationSourceHubspotUpdateAuthenticationPrivateApp(sourceHubspotUpdateAuthenticationPrivateApp SourceHubspotUpdateAuthenticationPrivateApp) SourceHubspotUpdateAuthentication {
	typ := SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateAuthenticationPrivateApp

	return SourceHubspotUpdateAuthentication{
		SourceHubspotUpdateAuthenticationPrivateApp: &sourceHubspotUpdateAuthenticationPrivateApp,
		Type: typ,
	}
}

func (u *SourceHubspotUpdateAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceHubspotUpdateAuthenticationOAuth := new(SourceHubspotUpdateAuthenticationOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHubspotUpdateAuthenticationOAuth); err == nil {
		u.SourceHubspotUpdateAuthenticationOAuth = sourceHubspotUpdateAuthenticationOAuth
		u.Type = SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateAuthenticationOAuth
		return nil
	}

	sourceHubspotUpdateAuthenticationPrivateApp := new(SourceHubspotUpdateAuthenticationPrivateApp)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceHubspotUpdateAuthenticationPrivateApp); err == nil {
		u.SourceHubspotUpdateAuthenticationPrivateApp = sourceHubspotUpdateAuthenticationPrivateApp
		u.Type = SourceHubspotUpdateAuthenticationTypeSourceHubspotUpdateAuthenticationPrivateApp
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceHubspotUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceHubspotUpdateAuthenticationOAuth != nil {
		return json.Marshal(u.SourceHubspotUpdateAuthenticationOAuth)
	}

	if u.SourceHubspotUpdateAuthenticationPrivateApp != nil {
		return json.Marshal(u.SourceHubspotUpdateAuthenticationPrivateApp)
	}

	return nil, nil
}

type SourceHubspotUpdate struct {
	// Choose how to authenticate to HubSpot.
	Credentials SourceHubspotUpdateAuthentication `json:"credentials"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
