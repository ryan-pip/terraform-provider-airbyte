// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceDynamodbResourceModel) ToCreateSDKType() *shared.SourceDynamodbCreateRequest {
	accessKeyID := r.Configuration.AccessKeyID.ValueString()
	endpoint := new(string)
	if !r.Configuration.Endpoint.IsUnknown() && !r.Configuration.Endpoint.IsNull() {
		*endpoint = r.Configuration.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	region := new(shared.SourceDynamodbDynamodbRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceDynamodbDynamodbRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	reservedAttributeNames := new(string)
	if !r.Configuration.ReservedAttributeNames.IsUnknown() && !r.Configuration.ReservedAttributeNames.IsNull() {
		*reservedAttributeNames = r.Configuration.ReservedAttributeNames.ValueString()
	} else {
		reservedAttributeNames = nil
	}
	secretAccessKey := r.Configuration.SecretAccessKey.ValueString()
	sourceType := shared.SourceDynamodbDynamodb(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceDynamodb{
		AccessKeyID:            accessKeyID,
		Endpoint:               endpoint,
		Region:                 region,
		ReservedAttributeNames: reservedAttributeNames,
		SecretAccessKey:        secretAccessKey,
		SourceType:             sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDynamodbCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDynamodbResourceModel) ToGetSDKType() *shared.SourceDynamodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDynamodbResourceModel) ToUpdateSDKType() *shared.SourceDynamodbPutRequest {
	accessKeyID := r.Configuration.AccessKeyID.ValueString()
	endpoint := new(string)
	if !r.Configuration.Endpoint.IsUnknown() && !r.Configuration.Endpoint.IsNull() {
		*endpoint = r.Configuration.Endpoint.ValueString()
	} else {
		endpoint = nil
	}
	region := new(shared.SourceDynamodbUpdateDynamodbRegion)
	if !r.Configuration.Region.IsUnknown() && !r.Configuration.Region.IsNull() {
		*region = shared.SourceDynamodbUpdateDynamodbRegion(r.Configuration.Region.ValueString())
	} else {
		region = nil
	}
	reservedAttributeNames := new(string)
	if !r.Configuration.ReservedAttributeNames.IsUnknown() && !r.Configuration.ReservedAttributeNames.IsNull() {
		*reservedAttributeNames = r.Configuration.ReservedAttributeNames.ValueString()
	} else {
		reservedAttributeNames = nil
	}
	secretAccessKey := r.Configuration.SecretAccessKey.ValueString()
	configuration := shared.SourceDynamodbUpdate{
		AccessKeyID:            accessKeyID,
		Endpoint:               endpoint,
		Region:                 region,
		ReservedAttributeNames: reservedAttributeNames,
		SecretAccessKey:        secretAccessKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceDynamodbPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceDynamodbResourceModel) ToDeleteSDKType() *shared.SourceDynamodbCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceDynamodbResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceDynamodbResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
