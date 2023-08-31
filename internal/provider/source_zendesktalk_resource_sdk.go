// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceZendeskTalkResourceModel) ToCreateSDKType() *shared.SourceZendeskTalkCreateRequest {
	var credentials *shared.SourceZendeskTalkAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskTalkAuthenticationAPIToken *shared.SourceZendeskTalkAuthenticationAPIToken
		if r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.APIToken.ValueString()
			authType := new(shared.SourceZendeskTalkAuthenticationAPITokenAuthType)
			if !r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.AuthType.IsNull() {
				*authType = shared.SourceZendeskTalkAuthenticationAPITokenAuthType(r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.AuthType.ValueString())
			} else {
				authType = nil
			}
			email := r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.Email.ValueString()
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskTalkAuthenticationAPIToken.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceZendeskTalkAuthenticationAPIToken = &shared.SourceZendeskTalkAuthenticationAPIToken{
				APIToken:             apiToken,
				AuthType:             authType,
				Email:                email,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceZendeskTalkAuthenticationAPIToken != nil {
			credentials = &shared.SourceZendeskTalkAuthentication{
				SourceZendeskTalkAuthenticationAPIToken: sourceZendeskTalkAuthenticationAPIToken,
			}
		}
		var sourceZendeskTalkAuthenticationOAuth20 *shared.SourceZendeskTalkAuthenticationOAuth20
		if r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AccessToken.ValueString()
			authType1 := new(shared.SourceZendeskTalkAuthenticationOAuth20AuthType)
			if !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AuthType.IsNull() {
				*authType1 = shared.SourceZendeskTalkAuthenticationOAuth20AuthType(r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskTalkAuthenticationOAuth20.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceZendeskTalkAuthenticationOAuth20 = &shared.SourceZendeskTalkAuthenticationOAuth20{
				AccessToken:          accessToken,
				AuthType:             authType1,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceZendeskTalkAuthenticationOAuth20 != nil {
			credentials = &shared.SourceZendeskTalkAuthentication{
				SourceZendeskTalkAuthenticationOAuth20: sourceZendeskTalkAuthenticationOAuth20,
			}
		}
	}
	sourceType := shared.SourceZendeskTalkZendeskTalk(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskTalk{
		Credentials: credentials,
		SourceType:  sourceType,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskTalkCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskTalkResourceModel) ToGetSDKType() *shared.SourceZendeskTalkCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskTalkResourceModel) ToUpdateSDKType() *shared.SourceZendeskTalkPutRequest {
	var credentials *shared.SourceZendeskTalkUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceZendeskTalkUpdateAuthenticationAPIToken *shared.SourceZendeskTalkUpdateAuthenticationAPIToken
		if r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.APIToken.ValueString()
			authType := new(shared.SourceZendeskTalkUpdateAuthenticationAPITokenAuthType)
			if !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.AuthType.IsNull() {
				*authType = shared.SourceZendeskTalkUpdateAuthenticationAPITokenAuthType(r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.AuthType.ValueString())
			} else {
				authType = nil
			}
			email := r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.Email.ValueString()
			var additionalProperties interface{}
			if !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationAPIToken.AdditionalProperties.ValueString()), &additionalProperties)
			}
			sourceZendeskTalkUpdateAuthenticationAPIToken = &shared.SourceZendeskTalkUpdateAuthenticationAPIToken{
				APIToken:             apiToken,
				AuthType:             authType,
				Email:                email,
				AdditionalProperties: additionalProperties,
			}
		}
		if sourceZendeskTalkUpdateAuthenticationAPIToken != nil {
			credentials = &shared.SourceZendeskTalkUpdateAuthentication{
				SourceZendeskTalkUpdateAuthenticationAPIToken: sourceZendeskTalkUpdateAuthenticationAPIToken,
			}
		}
		var sourceZendeskTalkUpdateAuthenticationOAuth20 *shared.SourceZendeskTalkUpdateAuthenticationOAuth20
		if r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AccessToken.ValueString()
			authType1 := new(shared.SourceZendeskTalkUpdateAuthenticationOAuth20AuthType)
			if !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AuthType.IsNull() {
				*authType1 = shared.SourceZendeskTalkUpdateAuthenticationOAuth20AuthType(r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			clientID := new(string)
			if !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.ClientID.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.ClientID.IsNull() {
				*clientID = r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.ClientID.ValueString()
			} else {
				clientID = nil
			}
			clientSecret := new(string)
			if !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.ClientSecret.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.ClientSecret.IsNull() {
				*clientSecret = r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.ClientSecret.ValueString()
			} else {
				clientSecret = nil
			}
			var additionalProperties1 interface{}
			if !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AdditionalProperties.IsUnknown() && !r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AdditionalProperties.IsNull() {
				_ = json.Unmarshal([]byte(r.Configuration.Credentials.SourceZendeskTalkUpdateAuthenticationOAuth20.AdditionalProperties.ValueString()), &additionalProperties1)
			}
			sourceZendeskTalkUpdateAuthenticationOAuth20 = &shared.SourceZendeskTalkUpdateAuthenticationOAuth20{
				AccessToken:          accessToken,
				AuthType:             authType1,
				ClientID:             clientID,
				ClientSecret:         clientSecret,
				AdditionalProperties: additionalProperties1,
			}
		}
		if sourceZendeskTalkUpdateAuthenticationOAuth20 != nil {
			credentials = &shared.SourceZendeskTalkUpdateAuthentication{
				SourceZendeskTalkUpdateAuthenticationOAuth20: sourceZendeskTalkUpdateAuthenticationOAuth20,
			}
		}
	}
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	subdomain := r.Configuration.Subdomain.ValueString()
	configuration := shared.SourceZendeskTalkUpdate{
		Credentials: credentials,
		StartDate:   startDate,
		Subdomain:   subdomain,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceZendeskTalkPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceZendeskTalkResourceModel) ToDeleteSDKType() *shared.SourceZendeskTalkCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceZendeskTalkResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceZendeskTalkResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
