// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceOracleConnectByServiceName struct {
	ConnectionType types.String `tfsdk:"connection_type"`
	ServiceName    types.String `tfsdk:"service_name"`
}
