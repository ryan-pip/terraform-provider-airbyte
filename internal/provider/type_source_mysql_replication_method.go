// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

type SourceMysqlReplicationMethod struct {
	SourceMysqlReplicationMethodLogicalReplicationCDC       *SourceMysqlReplicationMethodLogicalReplicationCDC `tfsdk:"source_mysql_replication_method_logical_replication_cdc"`
	SourceMysqlReplicationMethodStandard                    *SourceMssqlReplicationMethodStandard              `tfsdk:"source_mysql_replication_method_standard"`
	SourceMysqlUpdateReplicationMethodLogicalReplicationCDC *SourceMysqlReplicationMethodLogicalReplicationCDC `tfsdk:"source_mysql_update_replication_method_logical_replication_cdc"`
	SourceMysqlUpdateReplicationMethodStandard              *SourceMssqlReplicationMethodStandard              `tfsdk:"source_mysql_update_replication_method_standard"`
}
