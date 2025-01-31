// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// DestinationGoogleSheetsAuthenticationViaGoogleOAuth - Google API Credentials for connecting to Google Sheets and Google Drive APIs
type DestinationGoogleSheetsAuthenticationViaGoogleOAuth struct {
	// The Client ID of your Google Sheets developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Google Sheets developer application.
	ClientSecret string `json:"client_secret"`
	// The token for obtaining new access token.
	RefreshToken string `json:"refresh_token"`
}

type DestinationGoogleSheetsGoogleSheets string

const (
	DestinationGoogleSheetsGoogleSheetsGoogleSheets DestinationGoogleSheetsGoogleSheets = "google-sheets"
)

func (e DestinationGoogleSheetsGoogleSheets) ToPointer() *DestinationGoogleSheetsGoogleSheets {
	return &e
}

func (e *DestinationGoogleSheetsGoogleSheets) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-sheets":
		*e = DestinationGoogleSheetsGoogleSheets(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationGoogleSheetsGoogleSheets: %v", v)
	}
}

type DestinationGoogleSheets struct {
	// Google API Credentials for connecting to Google Sheets and Google Drive APIs
	Credentials     DestinationGoogleSheetsAuthenticationViaGoogleOAuth `json:"credentials"`
	DestinationType DestinationGoogleSheetsGoogleSheets                 `json:"destinationType"`
	// The link to your spreadsheet. See <a href='https://docs.airbyte.com/integrations/destinations/google-sheets#sheetlink'>this guide</a> for more details.
	SpreadsheetID string `json:"spreadsheet_id"`
}
