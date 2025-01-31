// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceGithubAuthenticationPersonalAccessTokenOptionTitle string

const (
	SourceGithubAuthenticationPersonalAccessTokenOptionTitlePatCredentials SourceGithubAuthenticationPersonalAccessTokenOptionTitle = "PAT Credentials"
)

func (e SourceGithubAuthenticationPersonalAccessTokenOptionTitle) ToPointer() *SourceGithubAuthenticationPersonalAccessTokenOptionTitle {
	return &e
}

func (e *SourceGithubAuthenticationPersonalAccessTokenOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "PAT Credentials":
		*e = SourceGithubAuthenticationPersonalAccessTokenOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGithubAuthenticationPersonalAccessTokenOptionTitle: %v", v)
	}
}

// SourceGithubAuthenticationPersonalAccessToken - Choose how to authenticate to GitHub
type SourceGithubAuthenticationPersonalAccessToken struct {
	OptionTitle *SourceGithubAuthenticationPersonalAccessTokenOptionTitle `json:"option_title,omitempty"`
	// Log into GitHub and then generate a <a href="https://github.com/settings/tokens">personal access token</a>. To load balance your API quota consumption across multiple API tokens, input multiple tokens separated with ","
	PersonalAccessToken string `json:"personal_access_token"`
}

type SourceGithubAuthenticationOAuthOptionTitle string

const (
	SourceGithubAuthenticationOAuthOptionTitleOAuthCredentials SourceGithubAuthenticationOAuthOptionTitle = "OAuth Credentials"
)

func (e SourceGithubAuthenticationOAuthOptionTitle) ToPointer() *SourceGithubAuthenticationOAuthOptionTitle {
	return &e
}

func (e *SourceGithubAuthenticationOAuthOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "OAuth Credentials":
		*e = SourceGithubAuthenticationOAuthOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGithubAuthenticationOAuthOptionTitle: %v", v)
	}
}

// SourceGithubAuthenticationOAuth - Choose how to authenticate to GitHub
type SourceGithubAuthenticationOAuth struct {
	// OAuth access token
	AccessToken string `json:"access_token"`
	// OAuth Client Id
	ClientID *string `json:"client_id,omitempty"`
	// OAuth Client secret
	ClientSecret *string                                     `json:"client_secret,omitempty"`
	OptionTitle  *SourceGithubAuthenticationOAuthOptionTitle `json:"option_title,omitempty"`
}

type SourceGithubAuthenticationType string

const (
	SourceGithubAuthenticationTypeSourceGithubAuthenticationOAuth               SourceGithubAuthenticationType = "source-github_Authentication_OAuth"
	SourceGithubAuthenticationTypeSourceGithubAuthenticationPersonalAccessToken SourceGithubAuthenticationType = "source-github_Authentication_Personal Access Token"
)

type SourceGithubAuthentication struct {
	SourceGithubAuthenticationOAuth               *SourceGithubAuthenticationOAuth
	SourceGithubAuthenticationPersonalAccessToken *SourceGithubAuthenticationPersonalAccessToken

	Type SourceGithubAuthenticationType
}

func CreateSourceGithubAuthenticationSourceGithubAuthenticationOAuth(sourceGithubAuthenticationOAuth SourceGithubAuthenticationOAuth) SourceGithubAuthentication {
	typ := SourceGithubAuthenticationTypeSourceGithubAuthenticationOAuth

	return SourceGithubAuthentication{
		SourceGithubAuthenticationOAuth: &sourceGithubAuthenticationOAuth,
		Type:                            typ,
	}
}

func CreateSourceGithubAuthenticationSourceGithubAuthenticationPersonalAccessToken(sourceGithubAuthenticationPersonalAccessToken SourceGithubAuthenticationPersonalAccessToken) SourceGithubAuthentication {
	typ := SourceGithubAuthenticationTypeSourceGithubAuthenticationPersonalAccessToken

	return SourceGithubAuthentication{
		SourceGithubAuthenticationPersonalAccessToken: &sourceGithubAuthenticationPersonalAccessToken,
		Type: typ,
	}
}

func (u *SourceGithubAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceGithubAuthenticationOAuth := new(SourceGithubAuthenticationOAuth)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGithubAuthenticationOAuth); err == nil {
		u.SourceGithubAuthenticationOAuth = sourceGithubAuthenticationOAuth
		u.Type = SourceGithubAuthenticationTypeSourceGithubAuthenticationOAuth
		return nil
	}

	sourceGithubAuthenticationPersonalAccessToken := new(SourceGithubAuthenticationPersonalAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceGithubAuthenticationPersonalAccessToken); err == nil {
		u.SourceGithubAuthenticationPersonalAccessToken = sourceGithubAuthenticationPersonalAccessToken
		u.Type = SourceGithubAuthenticationTypeSourceGithubAuthenticationPersonalAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceGithubAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceGithubAuthenticationOAuth != nil {
		return json.Marshal(u.SourceGithubAuthenticationOAuth)
	}

	if u.SourceGithubAuthenticationPersonalAccessToken != nil {
		return json.Marshal(u.SourceGithubAuthenticationPersonalAccessToken)
	}

	return nil, nil
}

type SourceGithubGithub string

const (
	SourceGithubGithubGithub SourceGithubGithub = "github"
)

func (e SourceGithubGithub) ToPointer() *SourceGithubGithub {
	return &e
}

func (e *SourceGithubGithub) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "github":
		*e = SourceGithubGithub(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGithubGithub: %v", v)
	}
}

type SourceGithub struct {
	// Space-delimited list of GitHub repository branches to pull commits for, e.g. `airbytehq/airbyte/master`. If no branches are specified for a repository, the default branch will be pulled.
	Branch *string `json:"branch,omitempty"`
	// Choose how to authenticate to GitHub
	Credentials *SourceGithubAuthentication `json:"credentials,omitempty"`
	// Space-delimited list of GitHub organizations/repositories, e.g. `airbytehq/airbyte` for single repository, `airbytehq/*` for get all repositories from organization and `airbytehq/airbyte airbytehq/another-repo` for multiple repositories.
	Repository string `json:"repository"`
	// The GitHub API allows for a maximum of 5000 requests per hour (15000 for Github Enterprise). You can specify a lower value to limit your use of the API quota.
	RequestsPerHour *int64             `json:"requests_per_hour,omitempty"`
	SourceType      SourceGithubGithub `json:"sourceType"`
	// The date from which you'd like to replicate data from GitHub in the format YYYY-MM-DDT00:00:00Z. For the streams which support this configuration, only data generated on or after the start date will be replicated. This field doesn't apply to all streams, see the <a href="https://docs.airbyte.com/integrations/sources/github">docs</a> for more info
	StartDate time.Time `json:"start_date"`
}
