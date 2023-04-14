// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceWoocommerce struct {
	APIKey     types.String `tfsdk:"api_key"`
	APISecret  types.String `tfsdk:"api_secret"`
	Shop       types.String `tfsdk:"shop"`
	SourceType types.String `tfsdk:"source_type"`
	StartDate  types.String `tfsdk:"start_date"`
}
