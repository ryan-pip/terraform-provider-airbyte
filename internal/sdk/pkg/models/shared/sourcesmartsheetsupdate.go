// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType string

const (
	SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthTypeAccessToken SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType = "access_token"
)

func (e SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType) ToPointer() *SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType {
	return &e
}

func (e *SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType: %v", v)
	}
}

type SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken struct {
	// The access token to use for accessing your data from Smartsheets. This access token must be generated by a user with at least read access to the data you'd like to replicate. Generate an access token in the Smartsheets main menu by clicking Account > Apps & Integrations > API Access. See the <a href="https://docs.airbyte.com/integrations/sources/smartsheets/#setup-guide">setup guide</a> for information on how to obtain this token.
	AccessToken string                                                            `json:"access_token"`
	AuthType    *SourceSmartsheetsUpdateAuthorizationMethodAPIAccessTokenAuthType `json:"auth_type,omitempty"`
}

type SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType string

const (
	SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthTypeOauth20 SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType = "oauth2.0"
)

func (e SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType) ToPointer() *SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType {
	return &e
}

func (e *SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType: %v", v)
	}
}

type SourceSmartsheetsUpdateAuthorizationMethodOAuth20 struct {
	// Access Token for making authenticated requests.
	AccessToken string                                                     `json:"access_token"`
	AuthType    *SourceSmartsheetsUpdateAuthorizationMethodOAuth20AuthType `json:"auth_type,omitempty"`
	// The API ID of the SmartSheets developer application.
	ClientID string `json:"client_id"`
	// The API Secret the SmartSheets developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access_token.
	RefreshToken string `json:"refresh_token"`
	// The date-time when the access token should be refreshed.
	TokenExpiryDate time.Time `json:"token_expiry_date"`
}

type SourceSmartsheetsUpdateAuthorizationMethodType string

const (
	SourceSmartsheetsUpdateAuthorizationMethodTypeSourceSmartsheetsUpdateAuthorizationMethodOAuth20        SourceSmartsheetsUpdateAuthorizationMethodType = "source-smartsheets-update_Authorization Method_OAuth2.0"
	SourceSmartsheetsUpdateAuthorizationMethodTypeSourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken SourceSmartsheetsUpdateAuthorizationMethodType = "source-smartsheets-update_Authorization Method_API Access Token"
)

type SourceSmartsheetsUpdateAuthorizationMethod struct {
	SourceSmartsheetsUpdateAuthorizationMethodOAuth20        *SourceSmartsheetsUpdateAuthorizationMethodOAuth20
	SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken *SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken

	Type SourceSmartsheetsUpdateAuthorizationMethodType
}

func CreateSourceSmartsheetsUpdateAuthorizationMethodSourceSmartsheetsUpdateAuthorizationMethodOAuth20(sourceSmartsheetsUpdateAuthorizationMethodOAuth20 SourceSmartsheetsUpdateAuthorizationMethodOAuth20) SourceSmartsheetsUpdateAuthorizationMethod {
	typ := SourceSmartsheetsUpdateAuthorizationMethodTypeSourceSmartsheetsUpdateAuthorizationMethodOAuth20

	return SourceSmartsheetsUpdateAuthorizationMethod{
		SourceSmartsheetsUpdateAuthorizationMethodOAuth20: &sourceSmartsheetsUpdateAuthorizationMethodOAuth20,
		Type: typ,
	}
}

func CreateSourceSmartsheetsUpdateAuthorizationMethodSourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken(sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken) SourceSmartsheetsUpdateAuthorizationMethod {
	typ := SourceSmartsheetsUpdateAuthorizationMethodTypeSourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken

	return SourceSmartsheetsUpdateAuthorizationMethod{
		SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken: &sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken,
		Type: typ,
	}
}

func (u *SourceSmartsheetsUpdateAuthorizationMethod) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceSmartsheetsUpdateAuthorizationMethodOAuth20 := new(SourceSmartsheetsUpdateAuthorizationMethodOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSmartsheetsUpdateAuthorizationMethodOAuth20); err == nil {
		u.SourceSmartsheetsUpdateAuthorizationMethodOAuth20 = sourceSmartsheetsUpdateAuthorizationMethodOAuth20
		u.Type = SourceSmartsheetsUpdateAuthorizationMethodTypeSourceSmartsheetsUpdateAuthorizationMethodOAuth20
		return nil
	}

	sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken := new(SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken); err == nil {
		u.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken = sourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken
		u.Type = SourceSmartsheetsUpdateAuthorizationMethodTypeSourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceSmartsheetsUpdateAuthorizationMethod) MarshalJSON() ([]byte, error) {
	if u.SourceSmartsheetsUpdateAuthorizationMethodOAuth20 != nil {
		return json.Marshal(u.SourceSmartsheetsUpdateAuthorizationMethodOAuth20)
	}

	if u.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken != nil {
		return json.Marshal(u.SourceSmartsheetsUpdateAuthorizationMethodAPIAccessToken)
	}

	return nil, nil
}

type SourceSmartsheetsUpdateValidenums string

const (
	SourceSmartsheetsUpdateValidenumsSheetcreatedAt   SourceSmartsheetsUpdateValidenums = "sheetcreatedAt"
	SourceSmartsheetsUpdateValidenumsSheetid          SourceSmartsheetsUpdateValidenums = "sheetid"
	SourceSmartsheetsUpdateValidenumsSheetmodifiedAt  SourceSmartsheetsUpdateValidenums = "sheetmodifiedAt"
	SourceSmartsheetsUpdateValidenumsSheetname        SourceSmartsheetsUpdateValidenums = "sheetname"
	SourceSmartsheetsUpdateValidenumsSheetpermalink   SourceSmartsheetsUpdateValidenums = "sheetpermalink"
	SourceSmartsheetsUpdateValidenumsSheetversion     SourceSmartsheetsUpdateValidenums = "sheetversion"
	SourceSmartsheetsUpdateValidenumsSheetaccessLevel SourceSmartsheetsUpdateValidenums = "sheetaccess_level"
	SourceSmartsheetsUpdateValidenumsRowID            SourceSmartsheetsUpdateValidenums = "row_id"
	SourceSmartsheetsUpdateValidenumsRowAccessLevel   SourceSmartsheetsUpdateValidenums = "row_access_level"
	SourceSmartsheetsUpdateValidenumsRowCreatedAt     SourceSmartsheetsUpdateValidenums = "row_created_at"
	SourceSmartsheetsUpdateValidenumsRowCreatedBy     SourceSmartsheetsUpdateValidenums = "row_created_by"
	SourceSmartsheetsUpdateValidenumsRowExpanded      SourceSmartsheetsUpdateValidenums = "row_expanded"
	SourceSmartsheetsUpdateValidenumsRowModifiedBy    SourceSmartsheetsUpdateValidenums = "row_modified_by"
	SourceSmartsheetsUpdateValidenumsRowParentID      SourceSmartsheetsUpdateValidenums = "row_parent_id"
	SourceSmartsheetsUpdateValidenumsRowPermalink     SourceSmartsheetsUpdateValidenums = "row_permalink"
	SourceSmartsheetsUpdateValidenumsRowNumber        SourceSmartsheetsUpdateValidenums = "row_number"
	SourceSmartsheetsUpdateValidenumsRowVersion       SourceSmartsheetsUpdateValidenums = "row_version"
)

func (e SourceSmartsheetsUpdateValidenums) ToPointer() *SourceSmartsheetsUpdateValidenums {
	return &e
}

func (e *SourceSmartsheetsUpdateValidenums) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sheetcreatedAt":
		fallthrough
	case "sheetid":
		fallthrough
	case "sheetmodifiedAt":
		fallthrough
	case "sheetname":
		fallthrough
	case "sheetpermalink":
		fallthrough
	case "sheetversion":
		fallthrough
	case "sheetaccess_level":
		fallthrough
	case "row_id":
		fallthrough
	case "row_access_level":
		fallthrough
	case "row_created_at":
		fallthrough
	case "row_created_by":
		fallthrough
	case "row_expanded":
		fallthrough
	case "row_modified_by":
		fallthrough
	case "row_parent_id":
		fallthrough
	case "row_permalink":
		fallthrough
	case "row_number":
		fallthrough
	case "row_version":
		*e = SourceSmartsheetsUpdateValidenums(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSmartsheetsUpdateValidenums: %v", v)
	}
}

type SourceSmartsheetsUpdate struct {
	Credentials SourceSmartsheetsUpdateAuthorizationMethod `json:"credentials"`
	// A List of available columns which metadata can be pulled from.
	MetadataFields []SourceSmartsheetsUpdateValidenums `json:"metadata_fields,omitempty"`
	// The spreadsheet ID. Find it by opening the spreadsheet then navigating to File > Properties
	SpreadsheetID string `json:"spreadsheet_id"`
	// Only rows modified after this date/time will be replicated. This should be an ISO 8601 string, for instance: `2000-01-01T13:00:00`
	StartDatetime *time.Time `json:"start_datetime,omitempty"`
}
