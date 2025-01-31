// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleSheets struct {
	Credentials     SourceGoogleSheetsAuthentication `tfsdk:"credentials"`
	NamesConversion types.Bool                       `tfsdk:"names_conversion"`
	RowBatchSize    types.Int64                      `tfsdk:"row_batch_size"`
	SourceType      types.String                     `tfsdk:"source_type"`
	SpreadsheetID   types.String                     `tfsdk:"spreadsheet_id"`
}
