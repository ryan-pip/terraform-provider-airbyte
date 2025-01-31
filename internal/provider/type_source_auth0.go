// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceAuth0 struct {
	BaseURL     types.String                    `tfsdk:"base_url"`
	Credentials SourceAuth0AuthenticationMethod `tfsdk:"credentials"`
	SourceType  types.String                    `tfsdk:"source_type"`
}
