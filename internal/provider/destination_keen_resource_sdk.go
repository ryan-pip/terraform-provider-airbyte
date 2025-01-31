// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationKeenResourceModel) ToCreateSDKType() *shared.DestinationKeenCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	destinationType := shared.DestinationKeenKeen(r.Configuration.DestinationType.ValueString())
	inferTimestamp := new(bool)
	if !r.Configuration.InferTimestamp.IsUnknown() && !r.Configuration.InferTimestamp.IsNull() {
		*inferTimestamp = r.Configuration.InferTimestamp.ValueBool()
	} else {
		inferTimestamp = nil
	}
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.DestinationKeen{
		APIKey:          apiKey,
		DestinationType: destinationType,
		InferTimestamp:  inferTimestamp,
		ProjectID:       projectID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationKeenCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationKeenResourceModel) ToGetSDKType() *shared.DestinationKeenCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationKeenResourceModel) ToUpdateSDKType() *shared.DestinationKeenPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	inferTimestamp := new(bool)
	if !r.Configuration.InferTimestamp.IsUnknown() && !r.Configuration.InferTimestamp.IsNull() {
		*inferTimestamp = r.Configuration.InferTimestamp.ValueBool()
	} else {
		inferTimestamp = nil
	}
	projectID := r.Configuration.ProjectID.ValueString()
	configuration := shared.DestinationKeenUpdate{
		APIKey:         apiKey,
		InferTimestamp: inferTimestamp,
		ProjectID:      projectID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationKeenPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationKeenResourceModel) ToDeleteSDKType() *shared.DestinationKeenCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationKeenResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationKeenResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
