// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceGoogleDirectory struct {
	CredentialsJSON types.String `tfsdk:"credentials_json"`
	Email           types.String `tfsdk:"email"`
	SourceType      types.String `tfsdk:"source_type"`
}
