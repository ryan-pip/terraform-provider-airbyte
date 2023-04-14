// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceTrello struct {
	BoardIds   []types.String `tfsdk:"board_ids"`
	Key        types.String   `tfsdk:"key"`
	SourceType types.String   `tfsdk:"source_type"`
	StartDate  types.String   `tfsdk:"start_date"`
	Token      types.String   `tfsdk:"token"`
}
