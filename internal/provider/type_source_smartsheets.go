// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceSmartsheets struct {
	Credentials   SourceSmartsheetsAuthorizationMethod `tfsdk:"credentials"`
	SourceType    types.String                         `tfsdk:"source_type"`
	SpreadsheetID types.String                         `tfsdk:"spreadsheet_id"`
	StartDatetime types.String                         `tfsdk:"start_datetime"`
}
