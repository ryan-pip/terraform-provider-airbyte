// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceNotion struct {
	Credentials *SourceNotionAuthenticateUsing `tfsdk:"credentials"`
	SourceType  types.String                   `tfsdk:"source_type"`
	StartDate   types.String                   `tfsdk:"start_date"`
}
