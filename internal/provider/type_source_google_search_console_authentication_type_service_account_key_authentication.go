// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleSearchConsoleAuthenticationTypeServiceAccountKeyAuthentication struct {
	AuthType           types.String `tfsdk:"auth_type"`
	Email              types.String `tfsdk:"email"`
	ServiceAccountInfo types.String `tfsdk:"service_account_info"`
}
