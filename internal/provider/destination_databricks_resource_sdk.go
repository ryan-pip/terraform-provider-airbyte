// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDatabricksResourceModel) ToCreateSDKType() *shared.DestinationDatabricksCreateRequest {
	acceptTerms := r.Configuration.AcceptTerms.ValueBool()
	var dataSource shared.DestinationDatabricksDataSource
	var destinationDatabricksDataSourceRecommendedManagedTables *shared.DestinationDatabricksDataSourceRecommendedManagedTables
	if r.Configuration.DataSource.DestinationDatabricksDataSourceRecommendedManagedTables != nil {
		dataSourceType := shared.DestinationDatabricksDataSourceRecommendedManagedTablesDataSourceType(r.Configuration.DataSource.DestinationDatabricksDataSourceRecommendedManagedTables.DataSourceType.ValueString())
		destinationDatabricksDataSourceRecommendedManagedTables = &shared.DestinationDatabricksDataSourceRecommendedManagedTables{
			DataSourceType: dataSourceType,
		}
	}
	if destinationDatabricksDataSourceRecommendedManagedTables != nil {
		dataSource = shared.DestinationDatabricksDataSource{
			DestinationDatabricksDataSourceRecommendedManagedTables: destinationDatabricksDataSourceRecommendedManagedTables,
		}
	}
	var destinationDatabricksDataSourceAmazonS3 *shared.DestinationDatabricksDataSourceAmazonS3
	if r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3 != nil {
		dataSourceType1 := shared.DestinationDatabricksDataSourceAmazonS3DataSourceType(r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.DataSourceType.ValueString())
		fileNamePattern := new(string)
		if !r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.FileNamePattern.IsUnknown() && !r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		s3AccessKeyID := r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.S3AccessKeyID.ValueString()
		s3BucketName := r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.S3BucketName.ValueString()
		s3BucketPath := r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.S3BucketPath.ValueString()
		s3BucketRegion := shared.DestinationDatabricksDataSourceAmazonS3S3BucketRegion(r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.S3BucketRegion.ValueString())
		s3SecretAccessKey := r.Configuration.DataSource.DestinationDatabricksDataSourceAmazonS3.S3SecretAccessKey.ValueString()
		destinationDatabricksDataSourceAmazonS3 = &shared.DestinationDatabricksDataSourceAmazonS3{
			DataSourceType:    dataSourceType1,
			FileNamePattern:   fileNamePattern,
			S3AccessKeyID:     s3AccessKeyID,
			S3BucketName:      s3BucketName,
			S3BucketPath:      s3BucketPath,
			S3BucketRegion:    s3BucketRegion,
			S3SecretAccessKey: s3SecretAccessKey,
		}
	}
	if destinationDatabricksDataSourceAmazonS3 != nil {
		dataSource = shared.DestinationDatabricksDataSource{
			DestinationDatabricksDataSourceAmazonS3: destinationDatabricksDataSourceAmazonS3,
		}
	}
	var destinationDatabricksDataSourceAzureBlobStorage *shared.DestinationDatabricksDataSourceAzureBlobStorage
	if r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage != nil {
		azureBlobStorageAccountName := r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.AzureBlobStorageAccountName.ValueString()
		azureBlobStorageContainerName := r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.AzureBlobStorageContainerName.ValueString()
		azureBlobStorageEndpointDomainName := new(string)
		if !r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.AzureBlobStorageEndpointDomainName.IsUnknown() && !r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.AzureBlobStorageEndpointDomainName.IsNull() {
			*azureBlobStorageEndpointDomainName = r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.AzureBlobStorageEndpointDomainName.ValueString()
		} else {
			azureBlobStorageEndpointDomainName = nil
		}
		azureBlobStorageSasToken := r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.AzureBlobStorageSasToken.ValueString()
		dataSourceType2 := shared.DestinationDatabricksDataSourceAzureBlobStorageDataSourceType(r.Configuration.DataSource.DestinationDatabricksDataSourceAzureBlobStorage.DataSourceType.ValueString())
		destinationDatabricksDataSourceAzureBlobStorage = &shared.DestinationDatabricksDataSourceAzureBlobStorage{
			AzureBlobStorageAccountName:        azureBlobStorageAccountName,
			AzureBlobStorageContainerName:      azureBlobStorageContainerName,
			AzureBlobStorageEndpointDomainName: azureBlobStorageEndpointDomainName,
			AzureBlobStorageSasToken:           azureBlobStorageSasToken,
			DataSourceType:                     dataSourceType2,
		}
	}
	if destinationDatabricksDataSourceAzureBlobStorage != nil {
		dataSource = shared.DestinationDatabricksDataSource{
			DestinationDatabricksDataSourceAzureBlobStorage: destinationDatabricksDataSourceAzureBlobStorage,
		}
	}
	database := new(string)
	if !r.Configuration.Database.IsUnknown() && !r.Configuration.Database.IsNull() {
		*database = r.Configuration.Database.ValueString()
	} else {
		database = nil
	}
	databricksHTTPPath := r.Configuration.DatabricksHTTPPath.ValueString()
	databricksPersonalAccessToken := r.Configuration.DatabricksPersonalAccessToken.ValueString()
	databricksPort := new(string)
	if !r.Configuration.DatabricksPort.IsUnknown() && !r.Configuration.DatabricksPort.IsNull() {
		*databricksPort = r.Configuration.DatabricksPort.ValueString()
	} else {
		databricksPort = nil
	}
	databricksServerHostname := r.Configuration.DatabricksServerHostname.ValueString()
	destinationType := shared.DestinationDatabricksDatabricks(r.Configuration.DestinationType.ValueString())
	enableSchemaEvolution := new(bool)
	if !r.Configuration.EnableSchemaEvolution.IsUnknown() && !r.Configuration.EnableSchemaEvolution.IsNull() {
		*enableSchemaEvolution = r.Configuration.EnableSchemaEvolution.ValueBool()
	} else {
		enableSchemaEvolution = nil
	}
	purgeStagingData := new(bool)
	if !r.Configuration.PurgeStagingData.IsUnknown() && !r.Configuration.PurgeStagingData.IsNull() {
		*purgeStagingData = r.Configuration.PurgeStagingData.ValueBool()
	} else {
		purgeStagingData = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	configuration := shared.DestinationDatabricks{
		AcceptTerms:                   acceptTerms,
		DataSource:                    dataSource,
		Database:                      database,
		DatabricksHTTPPath:            databricksHTTPPath,
		DatabricksPersonalAccessToken: databricksPersonalAccessToken,
		DatabricksPort:                databricksPort,
		DatabricksServerHostname:      databricksServerHostname,
		DestinationType:               destinationType,
		EnableSchemaEvolution:         enableSchemaEvolution,
		PurgeStagingData:              purgeStagingData,
		Schema:                        schema,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabricksCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabricksResourceModel) ToGetSDKType() *shared.DestinationDatabricksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabricksResourceModel) ToUpdateSDKType() *shared.DestinationDatabricksPutRequest {
	acceptTerms := r.Configuration.AcceptTerms.ValueBool()
	var dataSource shared.DestinationDatabricksUpdateDataSource
	var destinationDatabricksUpdateDataSourceRecommendedManagedTables *shared.DestinationDatabricksUpdateDataSourceRecommendedManagedTables
	if r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceRecommendedManagedTables != nil {
		dataSourceType := shared.DestinationDatabricksUpdateDataSourceRecommendedManagedTablesDataSourceType(r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceRecommendedManagedTables.DataSourceType.ValueString())
		destinationDatabricksUpdateDataSourceRecommendedManagedTables = &shared.DestinationDatabricksUpdateDataSourceRecommendedManagedTables{
			DataSourceType: dataSourceType,
		}
	}
	if destinationDatabricksUpdateDataSourceRecommendedManagedTables != nil {
		dataSource = shared.DestinationDatabricksUpdateDataSource{
			DestinationDatabricksUpdateDataSourceRecommendedManagedTables: destinationDatabricksUpdateDataSourceRecommendedManagedTables,
		}
	}
	var destinationDatabricksUpdateDataSourceAmazonS3 *shared.DestinationDatabricksUpdateDataSourceAmazonS3
	if r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3 != nil {
		dataSourceType1 := shared.DestinationDatabricksUpdateDataSourceAmazonS3DataSourceType(r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.DataSourceType.ValueString())
		fileNamePattern := new(string)
		if !r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.FileNamePattern.IsUnknown() && !r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.FileNamePattern.IsNull() {
			*fileNamePattern = r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.FileNamePattern.ValueString()
		} else {
			fileNamePattern = nil
		}
		s3AccessKeyID := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.S3AccessKeyID.ValueString()
		s3BucketName := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.S3BucketName.ValueString()
		s3BucketPath := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.S3BucketPath.ValueString()
		s3BucketRegion := shared.DestinationDatabricksUpdateDataSourceAmazonS3S3BucketRegion(r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.S3BucketRegion.ValueString())
		s3SecretAccessKey := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAmazonS3.S3SecretAccessKey.ValueString()
		destinationDatabricksUpdateDataSourceAmazonS3 = &shared.DestinationDatabricksUpdateDataSourceAmazonS3{
			DataSourceType:    dataSourceType1,
			FileNamePattern:   fileNamePattern,
			S3AccessKeyID:     s3AccessKeyID,
			S3BucketName:      s3BucketName,
			S3BucketPath:      s3BucketPath,
			S3BucketRegion:    s3BucketRegion,
			S3SecretAccessKey: s3SecretAccessKey,
		}
	}
	if destinationDatabricksUpdateDataSourceAmazonS3 != nil {
		dataSource = shared.DestinationDatabricksUpdateDataSource{
			DestinationDatabricksUpdateDataSourceAmazonS3: destinationDatabricksUpdateDataSourceAmazonS3,
		}
	}
	var destinationDatabricksUpdateDataSourceAzureBlobStorage *shared.DestinationDatabricksUpdateDataSourceAzureBlobStorage
	if r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage != nil {
		azureBlobStorageAccountName := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.AzureBlobStorageAccountName.ValueString()
		azureBlobStorageContainerName := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.AzureBlobStorageContainerName.ValueString()
		azureBlobStorageEndpointDomainName := new(string)
		if !r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.AzureBlobStorageEndpointDomainName.IsUnknown() && !r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.AzureBlobStorageEndpointDomainName.IsNull() {
			*azureBlobStorageEndpointDomainName = r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.AzureBlobStorageEndpointDomainName.ValueString()
		} else {
			azureBlobStorageEndpointDomainName = nil
		}
		azureBlobStorageSasToken := r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.AzureBlobStorageSasToken.ValueString()
		dataSourceType2 := shared.DestinationDatabricksUpdateDataSourceAzureBlobStorageDataSourceType(r.Configuration.DataSource.DestinationDatabricksUpdateDataSourceAzureBlobStorage.DataSourceType.ValueString())
		destinationDatabricksUpdateDataSourceAzureBlobStorage = &shared.DestinationDatabricksUpdateDataSourceAzureBlobStorage{
			AzureBlobStorageAccountName:        azureBlobStorageAccountName,
			AzureBlobStorageContainerName:      azureBlobStorageContainerName,
			AzureBlobStorageEndpointDomainName: azureBlobStorageEndpointDomainName,
			AzureBlobStorageSasToken:           azureBlobStorageSasToken,
			DataSourceType:                     dataSourceType2,
		}
	}
	if destinationDatabricksUpdateDataSourceAzureBlobStorage != nil {
		dataSource = shared.DestinationDatabricksUpdateDataSource{
			DestinationDatabricksUpdateDataSourceAzureBlobStorage: destinationDatabricksUpdateDataSourceAzureBlobStorage,
		}
	}
	database := new(string)
	if !r.Configuration.Database.IsUnknown() && !r.Configuration.Database.IsNull() {
		*database = r.Configuration.Database.ValueString()
	} else {
		database = nil
	}
	databricksHTTPPath := r.Configuration.DatabricksHTTPPath.ValueString()
	databricksPersonalAccessToken := r.Configuration.DatabricksPersonalAccessToken.ValueString()
	databricksPort := new(string)
	if !r.Configuration.DatabricksPort.IsUnknown() && !r.Configuration.DatabricksPort.IsNull() {
		*databricksPort = r.Configuration.DatabricksPort.ValueString()
	} else {
		databricksPort = nil
	}
	databricksServerHostname := r.Configuration.DatabricksServerHostname.ValueString()
	enableSchemaEvolution := new(bool)
	if !r.Configuration.EnableSchemaEvolution.IsUnknown() && !r.Configuration.EnableSchemaEvolution.IsNull() {
		*enableSchemaEvolution = r.Configuration.EnableSchemaEvolution.ValueBool()
	} else {
		enableSchemaEvolution = nil
	}
	purgeStagingData := new(bool)
	if !r.Configuration.PurgeStagingData.IsUnknown() && !r.Configuration.PurgeStagingData.IsNull() {
		*purgeStagingData = r.Configuration.PurgeStagingData.ValueBool()
	} else {
		purgeStagingData = nil
	}
	schema := new(string)
	if !r.Configuration.Schema.IsUnknown() && !r.Configuration.Schema.IsNull() {
		*schema = r.Configuration.Schema.ValueString()
	} else {
		schema = nil
	}
	configuration := shared.DestinationDatabricksUpdate{
		AcceptTerms:                   acceptTerms,
		DataSource:                    dataSource,
		Database:                      database,
		DatabricksHTTPPath:            databricksHTTPPath,
		DatabricksPersonalAccessToken: databricksPersonalAccessToken,
		DatabricksPort:                databricksPort,
		DatabricksServerHostname:      databricksServerHostname,
		EnableSchemaEvolution:         enableSchemaEvolution,
		PurgeStagingData:              purgeStagingData,
		Schema:                        schema,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDatabricksPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDatabricksResourceModel) ToDeleteSDKType() *shared.DestinationDatabricksCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDatabricksResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDatabricksResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
