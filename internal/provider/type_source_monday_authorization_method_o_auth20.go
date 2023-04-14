// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMondayAuthorizationMethodOAuth20 struct {
	AccessToken  types.String `tfsdk:"access_token"`
	AuthType     types.String `tfsdk:"auth_type"`
	ClientID     types.String `tfsdk:"client_id"`
	ClientSecret types.String `tfsdk:"client_secret"`
	Subdomain    types.String `tfsdk:"subdomain"`
}
