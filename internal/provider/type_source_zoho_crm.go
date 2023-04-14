// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceZohoCrm struct {
	ClientID      types.String `tfsdk:"client_id"`
	ClientSecret  types.String `tfsdk:"client_secret"`
	DcRegion      types.String `tfsdk:"dc_region"`
	Edition       types.String `tfsdk:"edition"`
	Environment   types.String `tfsdk:"environment"`
	RefreshToken  types.String `tfsdk:"refresh_token"`
	SourceType    types.String `tfsdk:"source_type"`
	StartDatetime types.String `tfsdk:"start_datetime"`
}
