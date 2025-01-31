// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourcePostgres struct {
	Database          types.String                     `tfsdk:"database"`
	Host              types.String                     `tfsdk:"host"`
	JdbcURLParams     types.String                     `tfsdk:"jdbc_url_params"`
	Password          types.String                     `tfsdk:"password"`
	Port              types.Int64                      `tfsdk:"port"`
	ReplicationMethod *SourcePostgresReplicationMethod `tfsdk:"replication_method"`
	Schemas           []types.String                   `tfsdk:"schemas"`
	SourceType        types.String                     `tfsdk:"source_type"`
	SslMode           *SourcePostgresSSLModes          `tfsdk:"ssl_mode"`
	TunnelMethod      *SourcePostgresSSHTunnelMethod   `tfsdk:"tunnel_method"`
	Username          types.String                     `tfsdk:"username"`
}
