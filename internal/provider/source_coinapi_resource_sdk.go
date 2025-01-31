// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceCoinAPIResourceModel) ToCreateSDKType() *shared.SourceCoinAPICreateRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	environment := shared.SourceCoinAPIEnvironment(r.Configuration.Environment.ValueString())
	limit := new(int64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueInt64()
	} else {
		limit = nil
	}
	period := r.Configuration.Period.ValueString()
	sourceType := shared.SourceCoinAPICoinAPI(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	symbolID := r.Configuration.SymbolID.ValueString()
	configuration := shared.SourceCoinAPI{
		APIKey:      apiKey,
		EndDate:     endDate,
		Environment: environment,
		Limit:       limit,
		Period:      period,
		SourceType:  sourceType,
		StartDate:   startDate,
		SymbolID:    symbolID,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceCoinAPICreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceCoinAPIResourceModel) ToGetSDKType() *shared.SourceCoinAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceCoinAPIResourceModel) ToUpdateSDKType() *shared.SourceCoinAPIPutRequest {
	apiKey := r.Configuration.APIKey.ValueString()
	endDate := new(string)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate = r.Configuration.EndDate.ValueString()
	} else {
		endDate = nil
	}
	environment := shared.SourceCoinAPIUpdateEnvironment(r.Configuration.Environment.ValueString())
	limit := new(int64)
	if !r.Configuration.Limit.IsUnknown() && !r.Configuration.Limit.IsNull() {
		*limit = r.Configuration.Limit.ValueInt64()
	} else {
		limit = nil
	}
	period := r.Configuration.Period.ValueString()
	startDate := r.Configuration.StartDate.ValueString()
	symbolID := r.Configuration.SymbolID.ValueString()
	configuration := shared.SourceCoinAPIUpdate{
		APIKey:      apiKey,
		EndDate:     endDate,
		Environment: environment,
		Limit:       limit,
		Period:      period,
		StartDate:   startDate,
		SymbolID:    symbolID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceCoinAPIPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceCoinAPIResourceModel) ToDeleteSDKType() *shared.SourceCoinAPICreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceCoinAPIResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceCoinAPIResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
