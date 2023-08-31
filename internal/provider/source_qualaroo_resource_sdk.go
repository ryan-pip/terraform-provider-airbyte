// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceQualarooResourceModel) ToCreateSDKType() *shared.SourceQualarooCreateRequest {
	key := r.Configuration.Key.ValueString()
	sourceType := shared.SourceQualarooQualaroo(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	var surveyIds []string = nil
	for _, surveyIdsItem := range r.Configuration.SurveyIds {
		surveyIds = append(surveyIds, surveyIdsItem.ValueString())
	}
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceQualaroo{
		Key:        key,
		SourceType: sourceType,
		StartDate:  startDate,
		SurveyIds:  surveyIds,
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
	out := shared.SourceQualarooCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQualarooResourceModel) ToGetSDKType() *shared.SourceQualarooCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceQualarooResourceModel) ToUpdateSDKType() *shared.SourceQualarooPutRequest {
	key := r.Configuration.Key.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	var surveyIds []string = nil
	for _, surveyIdsItem := range r.Configuration.SurveyIds {
		surveyIds = append(surveyIds, surveyIdsItem.ValueString())
	}
	token := r.Configuration.Token.ValueString()
	configuration := shared.SourceQualarooUpdate{
		Key:       key,
		StartDate: startDate,
		SurveyIds: surveyIds,
		Token:     token,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceQualarooPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceQualarooResourceModel) ToDeleteSDKType() *shared.SourceQualarooCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceQualarooResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceQualarooResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
