// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTrustpilotResourceModel) ToCreateSDKType() *shared.SourceTrustpilotCreateRequest {
	var businessUnits []string = nil
	for _, businessUnitsItem := range r.Configuration.BusinessUnits {
		businessUnits = append(businessUnits, businessUnitsItem.ValueString())
	}
	var credentials shared.SourceTrustpilotAuthorizationMethod
	var sourceTrustpilotAuthorizationMethodOAuth20 *shared.SourceTrustpilotAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceTrustpilotAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceTrustpilotAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceTrustpilotAuthorizationMethodOAuth20 = &shared.SourceTrustpilotAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceTrustpilotAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceTrustpilotAuthorizationMethod{
			SourceTrustpilotAuthorizationMethodOAuth20: sourceTrustpilotAuthorizationMethodOAuth20,
		}
	}
	var sourceTrustpilotAuthorizationMethodAPIKey *shared.SourceTrustpilotAuthorizationMethodAPIKey
	if r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodAPIKey != nil {
		authType1 := new(shared.SourceTrustpilotAuthorizationMethodAPIKeyAuthType)
		if !r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodAPIKey.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodAPIKey.AuthType.IsNull() {
			*authType1 = shared.SourceTrustpilotAuthorizationMethodAPIKeyAuthType(r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodAPIKey.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		clientId1 := r.Configuration.Credentials.SourceTrustpilotAuthorizationMethodAPIKey.ClientID.ValueString()
		sourceTrustpilotAuthorizationMethodAPIKey = &shared.SourceTrustpilotAuthorizationMethodAPIKey{
			AuthType: authType1,
			ClientID: clientId1,
		}
	}
	if sourceTrustpilotAuthorizationMethodAPIKey != nil {
		credentials = shared.SourceTrustpilotAuthorizationMethod{
			SourceTrustpilotAuthorizationMethodAPIKey: sourceTrustpilotAuthorizationMethodAPIKey,
		}
	}
	sourceType := shared.SourceTrustpilotTrustpilot(r.Configuration.SourceType.ValueString())
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceTrustpilot{
		BusinessUnits: businessUnits,
		Credentials:   credentials,
		SourceType:    sourceType,
		StartDate:     startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTrustpilotCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrustpilotResourceModel) ToGetSDKType() *shared.SourceTrustpilotCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrustpilotResourceModel) ToUpdateSDKType() *shared.SourceTrustpilotPutRequest {
	var businessUnits []string = nil
	for _, businessUnitsItem := range r.Configuration.BusinessUnits {
		businessUnits = append(businessUnits, businessUnitsItem.ValueString())
	}
	var credentials shared.SourceTrustpilotUpdateAuthorizationMethod
	var sourceTrustpilotUpdateAuthorizationMethodOAuth20 *shared.SourceTrustpilotUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceTrustpilotUpdateAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceTrustpilotUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceTrustpilotUpdateAuthorizationMethodOAuth20 = &shared.SourceTrustpilotUpdateAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceTrustpilotUpdateAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceTrustpilotUpdateAuthorizationMethod{
			SourceTrustpilotUpdateAuthorizationMethodOAuth20: sourceTrustpilotUpdateAuthorizationMethodOAuth20,
		}
	}
	var sourceTrustpilotUpdateAuthorizationMethodAPIKey *shared.SourceTrustpilotUpdateAuthorizationMethodAPIKey
	if r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodAPIKey != nil {
		authType1 := new(shared.SourceTrustpilotUpdateAuthorizationMethodAPIKeyAuthType)
		if !r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodAPIKey.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodAPIKey.AuthType.IsNull() {
			*authType1 = shared.SourceTrustpilotUpdateAuthorizationMethodAPIKeyAuthType(r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodAPIKey.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		clientId1 := r.Configuration.Credentials.SourceTrustpilotUpdateAuthorizationMethodAPIKey.ClientID.ValueString()
		sourceTrustpilotUpdateAuthorizationMethodAPIKey = &shared.SourceTrustpilotUpdateAuthorizationMethodAPIKey{
			AuthType: authType1,
			ClientID: clientId1,
		}
	}
	if sourceTrustpilotUpdateAuthorizationMethodAPIKey != nil {
		credentials = shared.SourceTrustpilotUpdateAuthorizationMethod{
			SourceTrustpilotUpdateAuthorizationMethodAPIKey: sourceTrustpilotUpdateAuthorizationMethodAPIKey,
		}
	}
	startDate := r.Configuration.StartDate.ValueString()
	configuration := shared.SourceTrustpilotUpdate{
		BusinessUnits: businessUnits,
		Credentials:   credentials,
		StartDate:     startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTrustpilotPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTrustpilotResourceModel) ToDeleteSDKType() *shared.SourceTrustpilotCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTrustpilotResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTrustpilotResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
