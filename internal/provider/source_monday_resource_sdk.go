// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceMondayResourceModel) ToCreateSDKType() *shared.SourceMondayCreateRequest {
	var credentials *shared.SourceMondayAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceMondayAuthorizationMethodOAuth20 *shared.SourceMondayAuthorizationMethodOAuth20
		if r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.AccessToken.ValueString()
			authType := shared.SourceMondayAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.AuthType.ValueString())
			clientID := r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.ClientSecret.ValueString()
			subdomain := new(string)
			if !r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.Subdomain.IsUnknown() && !r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.Subdomain.IsNull() {
				*subdomain = r.Configuration.Credentials.SourceMondayAuthorizationMethodOAuth20.Subdomain.ValueString()
			} else {
				subdomain = nil
			}
			sourceMondayAuthorizationMethodOAuth20 = &shared.SourceMondayAuthorizationMethodOAuth20{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Subdomain:    subdomain,
			}
		}
		if sourceMondayAuthorizationMethodOAuth20 != nil {
			credentials = &shared.SourceMondayAuthorizationMethod{
				SourceMondayAuthorizationMethodOAuth20: sourceMondayAuthorizationMethodOAuth20,
			}
		}
		var sourceMondayAuthorizationMethodAPIToken *shared.SourceMondayAuthorizationMethodAPIToken
		if r.Configuration.Credentials.SourceMondayAuthorizationMethodAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceMondayAuthorizationMethodAPIToken.APIToken.ValueString()
			authType1 := shared.SourceMondayAuthorizationMethodAPITokenAuthType(r.Configuration.Credentials.SourceMondayAuthorizationMethodAPIToken.AuthType.ValueString())
			sourceMondayAuthorizationMethodAPIToken = &shared.SourceMondayAuthorizationMethodAPIToken{
				APIToken: apiToken,
				AuthType: authType1,
			}
		}
		if sourceMondayAuthorizationMethodAPIToken != nil {
			credentials = &shared.SourceMondayAuthorizationMethod{
				SourceMondayAuthorizationMethodAPIToken: sourceMondayAuthorizationMethodAPIToken,
			}
		}
	}
	sourceType := shared.SourceMondayMonday(r.Configuration.SourceType.ValueString())
	configuration := shared.SourceMonday{
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
	out := shared.SourceMondayCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMondayResourceModel) ToGetSDKType() *shared.SourceMondayCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMondayResourceModel) ToUpdateSDKType() *shared.SourceMondayPutRequest {
	var credentials *shared.SourceMondayUpdateAuthorizationMethod
	if r.Configuration.Credentials != nil {
		var sourceMondayUpdateAuthorizationMethodOAuth20 *shared.SourceMondayUpdateAuthorizationMethodOAuth20
		if r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
			authType := shared.SourceMondayUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
			clientID := r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
			subdomain := new(string)
			if !r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.Subdomain.IsUnknown() && !r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.Subdomain.IsNull() {
				*subdomain = r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodOAuth20.Subdomain.ValueString()
			} else {
				subdomain = nil
			}
			sourceMondayUpdateAuthorizationMethodOAuth20 = &shared.SourceMondayUpdateAuthorizationMethodOAuth20{
				AccessToken:  accessToken,
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				Subdomain:    subdomain,
			}
		}
		if sourceMondayUpdateAuthorizationMethodOAuth20 != nil {
			credentials = &shared.SourceMondayUpdateAuthorizationMethod{
				SourceMondayUpdateAuthorizationMethodOAuth20: sourceMondayUpdateAuthorizationMethodOAuth20,
			}
		}
		var sourceMondayUpdateAuthorizationMethodAPIToken *shared.SourceMondayUpdateAuthorizationMethodAPIToken
		if r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodAPIToken.APIToken.ValueString()
			authType1 := shared.SourceMondayUpdateAuthorizationMethodAPITokenAuthType(r.Configuration.Credentials.SourceMondayUpdateAuthorizationMethodAPIToken.AuthType.ValueString())
			sourceMondayUpdateAuthorizationMethodAPIToken = &shared.SourceMondayUpdateAuthorizationMethodAPIToken{
				APIToken: apiToken,
				AuthType: authType1,
			}
		}
		if sourceMondayUpdateAuthorizationMethodAPIToken != nil {
			credentials = &shared.SourceMondayUpdateAuthorizationMethod{
				SourceMondayUpdateAuthorizationMethodAPIToken: sourceMondayUpdateAuthorizationMethodAPIToken,
			}
		}
	}
	configuration := shared.SourceMondayUpdate{
		Credentials: credentials,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceMondayPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceMondayResourceModel) ToDeleteSDKType() *shared.SourceMondayCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceMondayResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceMondayResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
