// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceZuoraResourceModel) ToCreateSDKType() *shared.SourceZuoraCreateRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dataQuery := shared.SourceZuoraDataQueryType(r.Configuration.DataQuery.ValueString())
	sourceType := shared.SourceZuoraZuora(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	tenantEndpoint := shared.SourceZuoraTenantEndpointLocation(r.Configuration.TenantEndpoint.ValueString())
	windowInDays := new(string)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueString()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceZuora{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		DataQuery:      dataQuery,
		SourceType:     sourceType,
		StartDate:      startDate,
		TenantEndpoint: tenantEndpoint,
		WindowInDays:   windowInDays,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZuoraCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZuoraResourceModel) ToGetSDKType() *shared.SourceZuoraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZuoraResourceModel) ToUpdateSDKType() *shared.SourceZuoraPutRequest {
	clientID := r.Configuration.ClientID.ValueString()
	clientSecret := r.Configuration.ClientSecret.ValueString()
	dataQuery := shared.SourceZuoraUpdateDataQueryType(r.Configuration.DataQuery.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	tenantEndpoint := shared.SourceZuoraUpdateTenantEndpointLocation(r.Configuration.TenantEndpoint.ValueString())
	windowInDays := new(string)
	if !r.Configuration.WindowInDays.IsUnknown() && !r.Configuration.WindowInDays.IsNull() {
		*windowInDays = r.Configuration.WindowInDays.ValueString()
	} else {
		windowInDays = nil
	}
	configuration := shared.SourceZuoraUpdate{
		ClientID:       clientID,
		ClientSecret:   clientSecret,
		DataQuery:      dataQuery,
		StartDate:      startDate,
		TenantEndpoint: tenantEndpoint,
		WindowInDays:   windowInDays,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZuoraPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZuoraResourceModel) ToDeleteSDKType() *shared.SourceZuoraCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZuoraResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZuoraResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
