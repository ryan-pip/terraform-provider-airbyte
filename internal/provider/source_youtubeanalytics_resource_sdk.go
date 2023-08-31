// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceYoutubeAnalyticsResourceModel) ToCreateSDKType() *shared.SourceYoutubeAnalyticsCreateRequest {
	clientID := r.Configuration.Credentials.ClientID.ValueString()
	clientSecret := r.Configuration.Credentials.ClientSecret.ValueString()
	refreshToken := r.Configuration.Credentials.RefreshToken.ValueString()
	var additionalProperties interface{}
	if !r.Configuration.Credentials.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AdditionalProperties.IsNull() {
		_ = json.Unmarshal([]byte(r.Configuration.Credentials.AdditionalProperties.ValueString()), &additionalProperties)
	}
	credentials := shared.SourceYoutubeAnalyticsAuthenticateViaOAuth20{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		RefreshToken:         refreshToken,
		AdditionalProperties: additionalProperties,
	}
	sourceType := shared.SourceYoutubeAnalyticsYoutubeAnalytics(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceYoutubeAnalytics{
		Credentials: credentials,
		SourceType:  sourceType,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceYoutubeAnalyticsCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceYoutubeAnalyticsResourceModel) ToGetSDKType() *shared.SourceYoutubeAnalyticsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceYoutubeAnalyticsResourceModel) ToUpdateSDKType() *shared.SourceYoutubeAnalyticsPutRequest {
	clientID := r.Configuration.Credentials.ClientID.ValueString()
	clientSecret := r.Configuration.Credentials.ClientSecret.ValueString()
	refreshToken := r.Configuration.Credentials.RefreshToken.ValueString()
	var additionalProperties interface{}
	if !r.Configuration.Credentials.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.AdditionalProperties.IsNull() {
		_ = json.Unmarshal([]byte(r.Configuration.Credentials.AdditionalProperties.ValueString()), &additionalProperties)
	}
	credentials := shared.SourceYoutubeAnalyticsUpdateAuthenticateViaOAuth20{
		ClientID:             clientID,
		ClientSecret:         clientSecret,
		RefreshToken:         refreshToken,
		AdditionalProperties: additionalProperties,
	}
	configuration := shared.SourceYoutubeAnalyticsUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceYoutubeAnalyticsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceYoutubeAnalyticsResourceModel) ToDeleteSDKType() *shared.SourceYoutubeAnalyticsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceYoutubeAnalyticsResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceYoutubeAnalyticsResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
