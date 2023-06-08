// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *WorkspaceResourceModel) ToCreateSDKType() *shared.WorkspaceCreateRequest {
	name := r.Name.ValueString()
	out := shared.WorkspaceCreateRequest{
		Name: name,
	}
	return &out
}

func (r *WorkspaceResourceModel) ToUpdateSDKType() *shared.WorkspaceUpdateRequest {
	name := r.Name.ValueString()
	out := shared.WorkspaceUpdateRequest{
		Name: name,
	}
	return &out
}

func (r *WorkspaceResourceModel) ToDeleteSDKType() *shared.WorkspaceCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *WorkspaceResourceModel) RefreshFromCreateResponse(resp *shared.WorkspaceResponse) {
	r.DataResidency = types.StringValue(string(resp.DataResidency))
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *WorkspaceResourceModel) RefreshFromUpdateResponse(resp *shared.WorkspaceResponse) {
	r.RefreshFromCreateResponse(resp)
}
