// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceRecreation struct {
	Apikey         types.String `tfsdk:"apikey"`
	QueryCampsites types.String `tfsdk:"query_campsites"`
	SourceType     types.String `tfsdk:"source_type"`
}
