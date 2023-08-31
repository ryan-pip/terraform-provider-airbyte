// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourcePrestashopResourceModel) ToCreateSDKType() *shared.SourcePrestashopCreateRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	sourceType := shared.SourcePrestashopPrestashop(r.Configuration.SourceType.ValueString())
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourcePrestashop{
		AccessKey:  accessKey,
		SourceType: sourceType,
		StartDate:  startDate,
		URL:        url,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePrestashopCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePrestashopResourceModel) ToGetSDKType() *shared.SourcePrestashopCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePrestashopResourceModel) ToUpdateSDKType() *shared.SourcePrestashopPutRequest {
	accessKey := r.Configuration.AccessKey.ValueString()
	startDate := customTypes.MustDateFromString(r.Configuration.StartDate.ValueString())
	url := r.Configuration.URL.ValueString()
	configuration := shared.SourcePrestashopUpdate{
		AccessKey: accessKey,
		StartDate: startDate,
		URL:       url,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourcePrestashopPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourcePrestashopResourceModel) ToDeleteSDKType() *shared.SourcePrestashopCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourcePrestashopResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourcePrestashopResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
