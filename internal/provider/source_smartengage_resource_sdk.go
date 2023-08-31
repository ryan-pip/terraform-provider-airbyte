// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSmartengageResourceModel) ToCreateSDKType() *shared.SourceSmartengageCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	sourceType := shared.SourceSmartengageSmartengage(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceSmartengage{
		APIKey:     apiKey,
		SourceType: sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmartengageCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartengageResourceModel) ToGetSDKType() *shared.SourceSmartengageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSmartengageResourceModel) ToUpdateSDKType() *shared.SourceSmartengagePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	configuration := shared.SourceSmartengageUpdate{
		APIKey: apiKey,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSmartengagePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSmartengageResourceModel) ToDeleteSDKType() *shared.SourceSmartengageCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSmartengageResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSmartengageResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
