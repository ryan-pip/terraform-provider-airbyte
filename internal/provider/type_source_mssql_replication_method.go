// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceMssqlReplicationMethod struct {
	SourceMssqlReplicationMethodLogicalReplicationCDC       *SourceMssqlReplicationMethodLogicalReplicationCDC `tfsdk:"source_mssql_replication_method_logical_replication_cdc"`
	SourceMssqlReplicationMethodStandard                    *SourceMssqlReplicationMethodStandard              `tfsdk:"source_mssql_replication_method_standard"`
	SourceMssqlUpdateReplicationMethodLogicalReplicationCDC *SourceMssqlReplicationMethodLogicalReplicationCDC `tfsdk:"source_mssql_update_replication_method_logical_replication_cdc"`
	SourceMssqlUpdateReplicationMethodStandard              *SourceMssqlReplicationMethodStandard              `tfsdk:"source_mssql_update_replication_method_standard"`
}
