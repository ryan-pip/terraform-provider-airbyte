// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceNotionAuthenticateUsingAccessTokenAuthTypeEnum string

const (
	SourceNotionAuthenticateUsingAccessTokenAuthTypeEnumToken SourceNotionAuthenticateUsingAccessTokenAuthTypeEnum = "token"
)

func (e *SourceNotionAuthenticateUsingAccessTokenAuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "token":
		*e = SourceNotionAuthenticateUsingAccessTokenAuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNotionAuthenticateUsingAccessTokenAuthTypeEnum: %s", s)
	}
}

// SourceNotionAuthenticateUsingAccessToken - Pick an authentication method.
type SourceNotionAuthenticateUsingAccessToken struct {
	AuthType SourceNotionAuthenticateUsingAccessTokenAuthTypeEnum `json:"auth_type"`
	// Notion API access token, see the <a href="https://developers.notion.com/docs/authorization">docs</a> for more information on how to obtain this token.
	Token string `json:"token"`
}

type SourceNotionAuthenticateUsingOAuth20AuthTypeEnum string

const (
	SourceNotionAuthenticateUsingOAuth20AuthTypeEnumOAuth20 SourceNotionAuthenticateUsingOAuth20AuthTypeEnum = "OAuth2.0"
)

func (e *SourceNotionAuthenticateUsingOAuth20AuthTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "OAuth2.0":
		*e = SourceNotionAuthenticateUsingOAuth20AuthTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNotionAuthenticateUsingOAuth20AuthTypeEnum: %s", s)
	}
}

// SourceNotionAuthenticateUsingOAuth20 - Pick an authentication method.
type SourceNotionAuthenticateUsingOAuth20 struct {
	// Access Token is a token you received by complete the OauthWebFlow of Notion.
	AccessToken string                                           `json:"access_token"`
	AuthType    SourceNotionAuthenticateUsingOAuth20AuthTypeEnum `json:"auth_type"`
	// The ClientID of your Notion integration.
	ClientID string `json:"client_id"`
	// The ClientSecret of your Notion integration.
	ClientSecret string `json:"client_secret"`
}

type SourceNotionAuthenticateUsingType string

const (
	SourceNotionAuthenticateUsingTypeSourceNotionAuthenticateUsingOAuth20     SourceNotionAuthenticateUsingType = "source-notion_Authenticate using_OAuth2.0"
	SourceNotionAuthenticateUsingTypeSourceNotionAuthenticateUsingAccessToken SourceNotionAuthenticateUsingType = "source-notion_Authenticate using_Access Token"
)

type SourceNotionAuthenticateUsing struct {
	SourceNotionAuthenticateUsingOAuth20     *SourceNotionAuthenticateUsingOAuth20
	SourceNotionAuthenticateUsingAccessToken *SourceNotionAuthenticateUsingAccessToken

	Type SourceNotionAuthenticateUsingType
}

func CreateSourceNotionAuthenticateUsingSourceNotionAuthenticateUsingOAuth20(sourceNotionAuthenticateUsingOAuth20 SourceNotionAuthenticateUsingOAuth20) SourceNotionAuthenticateUsing {
	typ := SourceNotionAuthenticateUsingTypeSourceNotionAuthenticateUsingOAuth20

	return SourceNotionAuthenticateUsing{
		SourceNotionAuthenticateUsingOAuth20: &sourceNotionAuthenticateUsingOAuth20,
		Type:                                 typ,
	}
}

func CreateSourceNotionAuthenticateUsingSourceNotionAuthenticateUsingAccessToken(sourceNotionAuthenticateUsingAccessToken SourceNotionAuthenticateUsingAccessToken) SourceNotionAuthenticateUsing {
	typ := SourceNotionAuthenticateUsingTypeSourceNotionAuthenticateUsingAccessToken

	return SourceNotionAuthenticateUsing{
		SourceNotionAuthenticateUsingAccessToken: &sourceNotionAuthenticateUsingAccessToken,
		Type:                                     typ,
	}
}

func (u *SourceNotionAuthenticateUsing) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceNotionAuthenticateUsingOAuth20 := new(SourceNotionAuthenticateUsingOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceNotionAuthenticateUsingOAuth20); err == nil {
		u.SourceNotionAuthenticateUsingOAuth20 = sourceNotionAuthenticateUsingOAuth20
		u.Type = SourceNotionAuthenticateUsingTypeSourceNotionAuthenticateUsingOAuth20
		return nil
	}

	sourceNotionAuthenticateUsingAccessToken := new(SourceNotionAuthenticateUsingAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceNotionAuthenticateUsingAccessToken); err == nil {
		u.SourceNotionAuthenticateUsingAccessToken = sourceNotionAuthenticateUsingAccessToken
		u.Type = SourceNotionAuthenticateUsingTypeSourceNotionAuthenticateUsingAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceNotionAuthenticateUsing) MarshalJSON() ([]byte, error) {
	if u.SourceNotionAuthenticateUsingOAuth20 != nil {
		return json.Marshal(u.SourceNotionAuthenticateUsingOAuth20)
	}

	if u.SourceNotionAuthenticateUsingAccessToken != nil {
		return json.Marshal(u.SourceNotionAuthenticateUsingAccessToken)
	}

	return nil, nil
}

type SourceNotionNotionEnum string

const (
	SourceNotionNotionEnumNotion SourceNotionNotionEnum = "notion"
)

func (e *SourceNotionNotionEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "notion":
		*e = SourceNotionNotionEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceNotionNotionEnum: %s", s)
	}
}

// SourceNotion - The values required to configure the source.
type SourceNotion struct {
	// Pick an authentication method.
	Credentials *SourceNotionAuthenticateUsing `json:"credentials,omitempty"`
	SourceType  SourceNotionNotionEnum         `json:"sourceType"`
	// UTC date and time in the format 2017-01-25T00:00:00.000Z. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
