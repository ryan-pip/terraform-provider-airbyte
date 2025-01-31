// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDynamodbResourceModel) ToCreateSDKType() *shared.DestinationDynamodbCreateRequest {
	accessKeyID := r.Configuration.AccessKeyID.ValueString()
	destinationType := shared.DestinationDynamodbDynamodb(r.Configuration.DestinationType.ValueString())
	dynamodbEndpoint := new(string)
	if !r.Configuration.DynamodbEndpoint.IsUnknown() && !r.Configuration.DynamodbEndpoint.IsNull() {
		*dynamodbEndpoint = r.Configuration.DynamodbEndpoint.ValueString()
	} else {
		dynamodbEndpoint = nil
	}
	dynamodbRegion := shared.DestinationDynamodbDynamoDBRegion(r.Configuration.DynamodbRegion.ValueString())
	dynamodbTableNamePrefix := r.Configuration.DynamodbTableNamePrefix.ValueString()
	secretAccessKey := r.Configuration.SecretAccessKey.ValueString()
	configuration := shared.DestinationDynamodb{
		AccessKeyID:             accessKeyID,
		DestinationType:         destinationType,
		DynamodbEndpoint:        dynamodbEndpoint,
		DynamodbRegion:          dynamodbRegion,
		DynamodbTableNamePrefix: dynamodbTableNamePrefix,
		SecretAccessKey:         secretAccessKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDynamodbCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDynamodbResourceModel) ToGetSDKType() *shared.DestinationDynamodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDynamodbResourceModel) ToUpdateSDKType() *shared.DestinationDynamodbPutRequest {
	accessKeyID := r.Configuration.AccessKeyID.ValueString()
	dynamodbEndpoint := new(string)
	if !r.Configuration.DynamodbEndpoint.IsUnknown() && !r.Configuration.DynamodbEndpoint.IsNull() {
		*dynamodbEndpoint = r.Configuration.DynamodbEndpoint.ValueString()
	} else {
		dynamodbEndpoint = nil
	}
	dynamodbRegion := shared.DestinationDynamodbUpdateDynamoDBRegion(r.Configuration.DynamodbRegion.ValueString())
	dynamodbTableNamePrefix := r.Configuration.DynamodbTableNamePrefix.ValueString()
	secretAccessKey := r.Configuration.SecretAccessKey.ValueString()
	configuration := shared.DestinationDynamodbUpdate{
		AccessKeyID:             accessKeyID,
		DynamodbEndpoint:        dynamodbEndpoint,
		DynamodbRegion:          dynamodbRegion,
		DynamodbTableNamePrefix: dynamodbTableNamePrefix,
		SecretAccessKey:         secretAccessKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDynamodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDynamodbResourceModel) ToDeleteSDKType() *shared.DestinationDynamodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDynamodbResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDynamodbResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
