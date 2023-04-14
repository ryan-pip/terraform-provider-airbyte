// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationAzureBlobStorage struct {
	AzureBlobStorageAccountKey         types.String                            `tfsdk:"azure_blob_storage_account_key"`
	AzureBlobStorageAccountName        types.String                            `tfsdk:"azure_blob_storage_account_name"`
	AzureBlobStorageContainerName      types.String                            `tfsdk:"azure_blob_storage_container_name"`
	AzureBlobStorageEndpointDomainName types.String                            `tfsdk:"azure_blob_storage_endpoint_domain_name"`
	AzureBlobStorageOutputBufferSize   types.Int64                             `tfsdk:"azure_blob_storage_output_buffer_size"`
	AzureBlobStorageSpillSize          types.Int64                             `tfsdk:"azure_blob_storage_spill_size"`
	DestinationType                    types.String                            `tfsdk:"destination_type"`
	Format                             DestinationAzureBlobStorageOutputFormat `tfsdk:"format"`
}
