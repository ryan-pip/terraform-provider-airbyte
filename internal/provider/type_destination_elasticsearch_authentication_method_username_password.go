// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationElasticsearchAuthenticationMethodUsernamePassword struct {
	Method   types.String `tfsdk:"method"`
	Password types.String `tfsdk:"password"`
	Username types.String `tfsdk:"username"`
}
