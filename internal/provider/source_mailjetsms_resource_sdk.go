// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMailjetSmsResourceModel) ToCreateSDKType() *shared.SourceMailjetSmsCreateRequest {
	endDate := new(int64)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueInt64()
	} else {
		endDate = nil
	}
	sourceType := shared.SourceMailjetSmsMailjetSms(r.Configuration.SourceType.ValueString())
	startDate := new(int64)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueInt64()
	} else {
		startDate = nil
	}
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceMailjetSms{
		EndDate:    endDate,
		SourceType: sourceType,
		StartDate:  startDate,
		Token:      token,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMailjetSmsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailjetSmsResourceModel) ToGetSDKType() *shared.SourceMailjetSmsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMailjetSmsResourceModel) ToUpdateSDKType() *shared.SourceMailjetSmsPutRequest {
	endDate := new(int64)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueInt64()
	} else {
		endDate = nil
	}
	startDate := new(int64)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate = r.Configuration.StartDate.ValueInt64()
	} else {
		startDate = nil
	}
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceMailjetSmsUpdate{
		EndDate:   endDate,
		StartDate: startDate,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMailjetSmsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMailjetSmsResourceModel) ToDeleteSDKType() *shared.SourceMailjetSmsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMailjetSmsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMailjetSmsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
