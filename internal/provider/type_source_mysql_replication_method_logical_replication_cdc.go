// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceMysqlReplicationMethodLogicalReplicationCDC struct {
	InitialWaitingSeconds types.Int64  `tfsdk:"initial_waiting_seconds"`
	Method                types.String `tfsdk:"method"`
	ServerTimeZone        types.String `tfsdk:"server_time_zone"`
}
