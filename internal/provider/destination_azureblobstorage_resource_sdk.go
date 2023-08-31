// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationAzureBlobStorageResourceModel) ToCreateSDKType() *shared.DestinationAzureBlobStorageCreateRequest {
	azureBlobStorageAccountKey := r.Configuration.AzureBlobStorageAccountKey.ValueString()
	azureBlobStorageAccountName := r.Configuration.AzureBlobStorageAccountName.ValueString()
	azureBlobStorageContainerName := new(string)
	if !r.Configuration.AzureBlobStorageContainerName.IsUnknown() && !r.Configuration.AzureBlobStorageContainerName.IsNull() {
		*azureBlobStorageContainerName = r.Configuration.AzureBlobStorageContainerName.ValueString()
	} else {
		azureBlobStorageContainerName = nil
	}
	azureBlobStorageEndpointDomainName := new(string)
	if !r.Configuration.AzureBlobStorageEndpointDomainName.IsUnknown() && !r.Configuration.AzureBlobStorageEndpointDomainName.IsNull() {
		*azureBlobStorageEndpointDomainName = r.Configuration.AzureBlobStorageEndpointDomainName.ValueString()
	} else {
		azureBlobStorageEndpointDomainName = nil
	}
	azureBlobStorageOutputBufferSize := new(int64)
	if !r.Configuration.AzureBlobStorageOutputBufferSize.IsUnknown() && !r.Configuration.AzureBlobStorageOutputBufferSize.IsNull() {
		*azureBlobStorageOutputBufferSize = r.Configuration.AzureBlobStorageOutputBufferSize.ValueInt64()
	} else {
		azureBlobStorageOutputBufferSize = nil
	}
	azureBlobStorageSpillSize := new(int64)
	if !r.Configuration.AzureBlobStorageSpillSize.IsUnknown() && !r.Configuration.AzureBlobStorageSpillSize.IsNull() {
		*azureBlobStorageSpillSize = r.Configuration.AzureBlobStorageSpillSize.ValueInt64()
	} else {
		azureBlobStorageSpillSize = nil
	}
	destinationType := shared.DestinationAzureBlobStorageAzureBlobStorage(r.Configuration.DestinationType.ValueString())
	var format shared.DestinationAzureBlobStorageOutputFormat
	var destinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues *shared.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues
	if r.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues != nil {
		flattening := shared.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValuesNormalizationFlattening(r.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues.Flattening.ValueString())
		formatType := shared.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValuesFormatType(r.Configuration.Format.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues.FormatType.ValueString())
		destinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues = &shared.DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues{
			Flattening: flattening,
			FormatType: formatType,
		}
	}
	if destinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues != nil {
		format = shared.DestinationAzureBlobStorageOutputFormat{
			DestinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues: destinationAzureBlobStorageOutputFormatCSVCommaSeparatedValues,
		}
	}
	var destinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON *shared.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON
	if r.Configuration.Format.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		formatType1 := shared.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSONFormatType(r.Configuration.Format.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON.FormatType.ValueString())
		destinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON = &shared.DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON{
			FormatType: formatType1,
		}
	}
	if destinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		format = shared.DestinationAzureBlobStorageOutputFormat{
			DestinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON: destinationAzureBlobStorageOutputFormatJSONLinesNewlineDelimitedJSON,
		}
	}
	configuration := shared.DestinationAzureBlobStorage{
		AzureBlobStorageAccountKey:         azureBlobStorageAccountKey,
		AzureBlobStorageAccountName:        azureBlobStorageAccountName,
		AzureBlobStorageContainerName:      azureBlobStorageContainerName,
		AzureBlobStorageEndpointDomainName: azureBlobStorageEndpointDomainName,
		AzureBlobStorageOutputBufferSize:   azureBlobStorageOutputBufferSize,
		AzureBlobStorageSpillSize:          azureBlobStorageSpillSize,
		DestinationType:                    destinationType,
		Format:                             format,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationAzureBlobStorageCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationAzureBlobStorageResourceModel) ToGetSDKType() *shared.DestinationAzureBlobStorageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationAzureBlobStorageResourceModel) ToUpdateSDKType() *shared.DestinationAzureBlobStoragePutRequest {
	azureBlobStorageAccountKey := r.Configuration.AzureBlobStorageAccountKey.ValueString()
	azureBlobStorageAccountName := r.Configuration.AzureBlobStorageAccountName.ValueString()
	azureBlobStorageContainerName := new(string)
	if !r.Configuration.AzureBlobStorageContainerName.IsUnknown() && !r.Configuration.AzureBlobStorageContainerName.IsNull() {
		*azureBlobStorageContainerName = r.Configuration.AzureBlobStorageContainerName.ValueString()
	} else {
		azureBlobStorageContainerName = nil
	}
	azureBlobStorageEndpointDomainName := new(string)
	if !r.Configuration.AzureBlobStorageEndpointDomainName.IsUnknown() && !r.Configuration.AzureBlobStorageEndpointDomainName.IsNull() {
		*azureBlobStorageEndpointDomainName = r.Configuration.AzureBlobStorageEndpointDomainName.ValueString()
	} else {
		azureBlobStorageEndpointDomainName = nil
	}
	azureBlobStorageOutputBufferSize := new(int64)
	if !r.Configuration.AzureBlobStorageOutputBufferSize.IsUnknown() && !r.Configuration.AzureBlobStorageOutputBufferSize.IsNull() {
		*azureBlobStorageOutputBufferSize = r.Configuration.AzureBlobStorageOutputBufferSize.ValueInt64()
	} else {
		azureBlobStorageOutputBufferSize = nil
	}
	azureBlobStorageSpillSize := new(int64)
	if !r.Configuration.AzureBlobStorageSpillSize.IsUnknown() && !r.Configuration.AzureBlobStorageSpillSize.IsNull() {
		*azureBlobStorageSpillSize = r.Configuration.AzureBlobStorageSpillSize.ValueInt64()
	} else {
		azureBlobStorageSpillSize = nil
	}
	var format shared.DestinationAzureBlobStorageUpdateOutputFormat
	var destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues *shared.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues
	if r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues != nil {
		flattening := shared.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesNormalizationFlattening(r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues.Flattening.ValueString())
		formatType := shared.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValuesFormatType(r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues.FormatType.ValueString())
		destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues = &shared.DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues{
			Flattening: flattening,
			FormatType: formatType,
		}
	}
	if destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues != nil {
		format = shared.DestinationAzureBlobStorageUpdateOutputFormat{
			DestinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues: destinationAzureBlobStorageUpdateOutputFormatCSVCommaSeparatedValues,
		}
	}
	var destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON *shared.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON
	if r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		formatType1 := shared.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSONFormatType(r.Configuration.Format.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON.FormatType.ValueString())
		destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON = &shared.DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON{
			FormatType: formatType1,
		}
	}
	if destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON != nil {
		format = shared.DestinationAzureBlobStorageUpdateOutputFormat{
			DestinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON: destinationAzureBlobStorageUpdateOutputFormatJSONLinesNewlineDelimitedJSON,
		}
	}
	configuration := shared.DestinationAzureBlobStorageUpdate{
		AzureBlobStorageAccountKey:         azureBlobStorageAccountKey,
		AzureBlobStorageAccountName:        azureBlobStorageAccountName,
		AzureBlobStorageContainerName:      azureBlobStorageContainerName,
		AzureBlobStorageEndpointDomainName: azureBlobStorageEndpointDomainName,
		AzureBlobStorageOutputBufferSize:   azureBlobStorageOutputBufferSize,
		AzureBlobStorageSpillSize:          azureBlobStorageSpillSize,
		Format:                             format,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationAzureBlobStoragePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationAzureBlobStorageResourceModel) ToDeleteSDKType() *shared.DestinationAzureBlobStorageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationAzureBlobStorageResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationAzureBlobStorageResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
