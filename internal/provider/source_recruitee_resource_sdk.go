// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceRecruiteeResourceModel) ToCreateSDKType() *shared.SourceRecruiteeCreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	companyID := r.Configuration.CompanyID.ValueInt64()
	sourceType := shared.SourceRecruiteeRecruitee(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceRecruitee{
		APIKey:     apiKey,
		CompanyID:  companyID,
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
	out := shared.SourceRecruiteeCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecruiteeResourceModel) ToGetSDKType() *shared.SourceRecruiteeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRecruiteeResourceModel) ToUpdateSDKType() *shared.SourceRecruiteePutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	companyID := r.Configuration.CompanyID.ValueInt64()
	configuration := shared.SourceRecruiteeUpdate{
		APIKey:    apiKey,
		CompanyID: companyID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceRecruiteePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceRecruiteeResourceModel) ToDeleteSDKType() *shared.SourceRecruiteeCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceRecruiteeResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceRecruiteeResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
