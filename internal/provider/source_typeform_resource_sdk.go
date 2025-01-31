// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceTypeformResourceModel) ToCreateSDKType() *shared.SourceTypeformCreateRequest {
	var credentials shared.SourceTypeformAuthorizationMethod
	var sourceTypeformAuthorizationMethodOAuth20 *shared.SourceTypeformAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceTypeformAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceTypeformAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceTypeformAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceTypeformAuthorizationMethodOAuth20 = &shared.SourceTypeformAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceTypeformAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceTypeformAuthorizationMethod{
			SourceTypeformAuthorizationMethodOAuth20: sourceTypeformAuthorizationMethodOAuth20,
		}
	}
	var sourceTypeformAuthorizationMethodPrivateToken *shared.SourceTypeformAuthorizationMethodPrivateToken
	if r.Configuration.Credentials.SourceTypeformAuthorizationMethodPrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceTypeformAuthorizationMethodPrivateToken.AccessToken.ValueString()
		authType1 := new(shared.SourceTypeformAuthorizationMethodPrivateTokenAuthType)
		if !r.Configuration.Credentials.SourceTypeformAuthorizationMethodPrivateToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTypeformAuthorizationMethodPrivateToken.AuthType.IsNull() {
			*authType1 = shared.SourceTypeformAuthorizationMethodPrivateTokenAuthType(r.Configuration.Credentials.SourceTypeformAuthorizationMethodPrivateToken.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		sourceTypeformAuthorizationMethodPrivateToken = &shared.SourceTypeformAuthorizationMethodPrivateToken{
			AccessToken: accessToken1,
			AuthType:    authType1,
		}
	}
	if sourceTypeformAuthorizationMethodPrivateToken != nil {
		credentials = shared.SourceTypeformAuthorizationMethod{
			SourceTypeformAuthorizationMethodPrivateToken: sourceTypeformAuthorizationMethodPrivateToken,
		}
	}
	var formIds []string = nil
	for _, formIdsItem := range r.Configuration.FormIds {
		formIds = append(formIds, formIdsItem.ValueString())
	}
	sourceType := shared.SourceTypeformTypeform(r.Configuration.SourceType.ValueString())
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceTypeform{
		Credentials: credentials,
		FormIds:     formIds,
		SourceType:  sourceType,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTypeformCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTypeformResourceModel) ToGetSDKType() *shared.SourceTypeformCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTypeformResourceModel) ToUpdateSDKType() *shared.SourceTypeformPutRequest {
	var credentials shared.SourceTypeformUpdateAuthorizationMethod
	var sourceTypeformUpdateAuthorizationMethodOAuth20 *shared.SourceTypeformUpdateAuthorizationMethodOAuth20
	if r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20 != nil {
		accessToken := r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.AccessToken.ValueString()
		authType := new(shared.SourceTypeformUpdateAuthorizationMethodOAuth20AuthType)
		if !r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.AuthType.IsNull() {
			*authType = shared.SourceTypeformUpdateAuthorizationMethodOAuth20AuthType(r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.AuthType.ValueString())
		} else {
			authType = nil
		}
		clientID := r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.ClientID.ValueString()
		clientSecret := r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.ClientSecret.ValueString()
		refreshToken := r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.RefreshToken.ValueString()
		tokenExpiryDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodOAuth20.TokenExpiryDate.ValueString())
		sourceTypeformUpdateAuthorizationMethodOAuth20 = &shared.SourceTypeformUpdateAuthorizationMethodOAuth20{
			AccessToken:     accessToken,
			AuthType:        authType,
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			RefreshToken:    refreshToken,
			TokenExpiryDate: tokenExpiryDate,
		}
	}
	if sourceTypeformUpdateAuthorizationMethodOAuth20 != nil {
		credentials = shared.SourceTypeformUpdateAuthorizationMethod{
			SourceTypeformUpdateAuthorizationMethodOAuth20: sourceTypeformUpdateAuthorizationMethodOAuth20,
		}
	}
	var sourceTypeformUpdateAuthorizationMethodPrivateToken *shared.SourceTypeformUpdateAuthorizationMethodPrivateToken
	if r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodPrivateToken != nil {
		accessToken1 := r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodPrivateToken.AccessToken.ValueString()
		authType1 := new(shared.SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType)
		if !r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodPrivateToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodPrivateToken.AuthType.IsNull() {
			*authType1 = shared.SourceTypeformUpdateAuthorizationMethodPrivateTokenAuthType(r.Configuration.Credentials.SourceTypeformUpdateAuthorizationMethodPrivateToken.AuthType.ValueString())
		} else {
			authType1 = nil
		}
		sourceTypeformUpdateAuthorizationMethodPrivateToken = &shared.SourceTypeformUpdateAuthorizationMethodPrivateToken{
			AccessToken: accessToken1,
			AuthType:    authType1,
		}
	}
	if sourceTypeformUpdateAuthorizationMethodPrivateToken != nil {
		credentials = shared.SourceTypeformUpdateAuthorizationMethod{
			SourceTypeformUpdateAuthorizationMethodPrivateToken: sourceTypeformUpdateAuthorizationMethodPrivateToken,
		}
	}
	var formIds []string = nil
	for _, formIdsItem := range r.Configuration.FormIds {
		formIds = append(formIds, formIdsItem.ValueString())
	}
	startDate := new(time.Time)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		*startDate, _ = time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceTypeformUpdate{
		Credentials: credentials,
		FormIds:     formIds,
		StartDate:   startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTypeformPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTypeformResourceModel) ToDeleteSDKType() *shared.SourceTypeformCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTypeformResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTypeformResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
