// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAmplitude struct {
	APIKey           types.String `tfsdk:"api_key"`
	DataRegion       types.String `tfsdk:"data_region"`
	RequestTimeRange types.Int64  `tfsdk:"request_time_range"`
	SecretKey        types.String `tfsdk:"secret_key"`
	SourceType       types.String `tfsdk:"source_type"`
	StartDate        types.String `tfsdk:"start_date"`
}
