// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceKlarna struct {
	Password   types.String `tfsdk:"password"`
	Playground types.Bool   `tfsdk:"playground"`
	Region     types.String `tfsdk:"region"`
	SourceType types.String `tfsdk:"source_type"`
	Username   types.String `tfsdk:"username"`
}
