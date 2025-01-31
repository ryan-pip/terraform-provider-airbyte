// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceSquareResourceModel) ToCreateSDKType() *shared.SourceSquareCreateRequest {
	var credentials *shared.SourceSquareAuthentication
	if r.Configuration.Credentials != nil {
		var sourceSquareAuthenticationOauthAuthentication *shared.SourceSquareAuthenticationOauthAuthentication
		if r.Configuration.Credentials.SourceSquareAuthenticationOauthAuthentication != nil {
			authType := shared.SourceSquareAuthenticationOauthAuthenticationAuthType(r.Configuration.Credentials.SourceSquareAuthenticationOauthAuthentication.AuthType.ValueString())
			clientID := r.Configuration.Credentials.SourceSquareAuthenticationOauthAuthentication.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceSquareAuthenticationOauthAuthentication.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceSquareAuthenticationOauthAuthentication.RefreshToken.ValueString()
			sourceSquareAuthenticationOauthAuthentication = &shared.SourceSquareAuthenticationOauthAuthentication{
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceSquareAuthenticationOauthAuthentication != nil {
			credentials = &shared.SourceSquareAuthentication{
				SourceSquareAuthenticationOauthAuthentication: sourceSquareAuthenticationOauthAuthentication,
			}
		}
		var sourceSquareAuthenticationAPIKey *shared.SourceSquareAuthenticationAPIKey
		if r.Configuration.Credentials.SourceSquareAuthenticationAPIKey != nil {
			apiKey := r.Configuration.Credentials.SourceSquareAuthenticationAPIKey.APIKey.ValueString()
			authType1 := shared.SourceSquareAuthenticationAPIKeyAuthType(r.Configuration.Credentials.SourceSquareAuthenticationAPIKey.AuthType.ValueString())
			sourceSquareAuthenticationAPIKey = &shared.SourceSquareAuthenticationAPIKey{
				APIKey:   apiKey,
				AuthType: authType1,
			}
		}
		if sourceSquareAuthenticationAPIKey != nil {
			credentials = &shared.SourceSquareAuthentication{
				SourceSquareAuthenticationAPIKey: sourceSquareAuthenticationAPIKey,
			}
		}
	}
	includeDeletedObjects := new(bool)
	if !r.Configuration.IncludeDeletedObjects.IsUnknown() && !r.Configuration.IncludeDeletedObjects.IsNull() {
		*includeDeletedObjects = r.Configuration.IncludeDeletedObjects.ValueBool()
	} else {
		includeDeletedObjects = nil
	}
	isSandbox := r.Configuration.IsSandbox.ValueBool()
	sourceType := shared.SourceSquareSquare(r.Configuration.SourceType.ValueString())
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceSquare{
		Credentials:           credentials,
		IncludeDeletedObjects: includeDeletedObjects,
		IsSandbox:             isSandbox,
		SourceType:            sourceType,
		StartDate:             startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSquareCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSquareResourceModel) ToGetSDKType() *shared.SourceSquareCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSquareResourceModel) ToUpdateSDKType() *shared.SourceSquarePutRequest {
	var credentials *shared.SourceSquareUpdateAuthentication
	if r.Configuration.Credentials != nil {
		var sourceSquareUpdateAuthenticationOauthAuthentication *shared.SourceSquareUpdateAuthenticationOauthAuthentication
		if r.Configuration.Credentials.SourceSquareUpdateAuthenticationOauthAuthentication != nil {
			authType := shared.SourceSquareUpdateAuthenticationOauthAuthenticationAuthType(r.Configuration.Credentials.SourceSquareUpdateAuthenticationOauthAuthentication.AuthType.ValueString())
			clientID := r.Configuration.Credentials.SourceSquareUpdateAuthenticationOauthAuthentication.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceSquareUpdateAuthenticationOauthAuthentication.ClientSecret.ValueString()
			refreshToken := r.Configuration.Credentials.SourceSquareUpdateAuthenticationOauthAuthentication.RefreshToken.ValueString()
			sourceSquareUpdateAuthenticationOauthAuthentication = &shared.SourceSquareUpdateAuthenticationOauthAuthentication{
				AuthType:     authType,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				RefreshToken: refreshToken,
			}
		}
		if sourceSquareUpdateAuthenticationOauthAuthentication != nil {
			credentials = &shared.SourceSquareUpdateAuthentication{
				SourceSquareUpdateAuthenticationOauthAuthentication: sourceSquareUpdateAuthenticationOauthAuthentication,
			}
		}
		var sourceSquareUpdateAuthenticationAPIKey *shared.SourceSquareUpdateAuthenticationAPIKey
		if r.Configuration.Credentials.SourceSquareUpdateAuthenticationAPIKey != nil {
			apiKey := r.Configuration.Credentials.SourceSquareUpdateAuthenticationAPIKey.APIKey.ValueString()
			authType1 := shared.SourceSquareUpdateAuthenticationAPIKeyAuthType(r.Configuration.Credentials.SourceSquareUpdateAuthenticationAPIKey.AuthType.ValueString())
			sourceSquareUpdateAuthenticationAPIKey = &shared.SourceSquareUpdateAuthenticationAPIKey{
				APIKey:   apiKey,
				AuthType: authType1,
			}
		}
		if sourceSquareUpdateAuthenticationAPIKey != nil {
			credentials = &shared.SourceSquareUpdateAuthentication{
				SourceSquareUpdateAuthenticationAPIKey: sourceSquareUpdateAuthenticationAPIKey,
			}
		}
	}
	includeDeletedObjects := new(bool)
	if !r.Configuration.IncludeDeletedObjects.IsUnknown() && !r.Configuration.IncludeDeletedObjects.IsNull() {
		*includeDeletedObjects = r.Configuration.IncludeDeletedObjects.ValueBool()
	} else {
		includeDeletedObjects = nil
	}
	isSandbox := r.Configuration.IsSandbox.ValueBool()
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceSquareUpdate{
		Credentials:           credentials,
		IncludeDeletedObjects: includeDeletedObjects,
		IsSandbox:             isSandbox,
		StartDate:             startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSquarePutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSquareResourceModel) ToDeleteSDKType() *shared.SourceSquareCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSquareResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSquareResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
