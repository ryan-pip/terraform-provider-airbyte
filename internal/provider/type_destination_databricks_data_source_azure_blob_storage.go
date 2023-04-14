// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationDatabricksDataSourceAzureBlobStorage struct {
	AzureBlobStorageAccountName        types.String `tfsdk:"azure_blob_storage_account_name"`
	AzureBlobStorageContainerName      types.String `tfsdk:"azure_blob_storage_container_name"`
	AzureBlobStorageEndpointDomainName types.String `tfsdk:"azure_blob_storage_endpoint_domain_name"`
	AzureBlobStorageSasToken           types.String `tfsdk:"azure_blob_storage_sas_token"`
	DataSourceType                     types.String `tfsdk:"data_source_type"`
}
