// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAircall struct {
	APIID      types.String `tfsdk:"api_id"`
	APIToken   types.String `tfsdk:"api_token"`
	SourceType types.String `tfsdk:"source_type"`
	StartDate  types.String `tfsdk:"start_date"`
}
