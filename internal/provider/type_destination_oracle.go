// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationOracle struct {
	DestinationType types.String                      `tfsdk:"destination_type"`
	Host            types.String                      `tfsdk:"host"`
	JdbcURLParams   types.String                      `tfsdk:"jdbc_url_params"`
	Password        types.String                      `tfsdk:"password"`
	Port            types.Int64                       `tfsdk:"port"`
	Schema          types.String                      `tfsdk:"schema"`
	Sid             types.String                      `tfsdk:"sid"`
	TunnelMethod    *DestinationOracleSSHTunnelMethod `tfsdk:"tunnel_method"`
	Username        types.String                      `tfsdk:"username"`
}
