// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceIp2whoisResourceModel) ToCreateSDKType() *shared.SourceIp2whoisCreateRequest {
	apiKey := new(string)
	if !r.Configuration.APIKey.IsUnknown() && !r.Configuration.APIKey.IsNull() {
		*apiKey = r.Configuration.APIKey.ValueString()
	} else {
		apiKey = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	sourceType := new(shared.SourceIp2whoisIp2whois)
	if !r.Configuration.SourceType.IsUnknown() && !r.Configuration.SourceType.IsNull() {
		*sourceType = shared.SourceIp2whoisIp2whois(r.Configuration.SourceType.ValueString())
	} else {
		sourceType = nil
	}
	configuration := shared.SourceIp2whois{
		APIKey:     apiKey,
		Domain:     domain,
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
	out := shared.SourceIp2whoisCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceIp2whoisResourceModel) ToGetSDKType() *shared.SourceIp2whoisCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceIp2whoisResourceModel) ToUpdateSDKType() *shared.SourceIp2whoisPutRequest {
	apiKey := new(string)
	if !r.Configuration.APIKey.IsUnknown() && !r.Configuration.APIKey.IsNull() {
		*apiKey = r.Configuration.APIKey.ValueString()
	} else {
		apiKey = nil
	}
	domain := new(string)
	if !r.Configuration.Domain.IsUnknown() && !r.Configuration.Domain.IsNull() {
		*domain = r.Configuration.Domain.ValueString()
	} else {
		domain = nil
	}
	configuration := shared.SourceIp2whoisUpdate{
		APIKey: apiKey,
		Domain: domain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceIp2whoisPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceIp2whoisResourceModel) ToDeleteSDKType() *shared.SourceIp2whoisCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceIp2whoisResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceIp2whoisResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
