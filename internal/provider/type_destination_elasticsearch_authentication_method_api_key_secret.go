// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationElasticsearchAuthenticationMethodAPIKeySecret struct {
	APIKeyID     types.String `tfsdk:"api_key_id"`
	APIKeySecret types.String `tfsdk:"api_key_secret"`
	Method       types.String `tfsdk:"method"`
}
