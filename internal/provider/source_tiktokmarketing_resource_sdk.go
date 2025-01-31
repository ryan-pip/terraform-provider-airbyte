// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	customTypes "github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceTiktokMarketingResourceModel) ToCreateSDKType() *shared.SourceTiktokMarketingCreateRequest {
	attributionWindow := new(int64)
	if !r.Configuration.AttributionWindow.IsUnknown() && !r.Configuration.AttributionWindow.IsNull() {
		*attributionWindow = r.Configuration.AttributionWindow.ValueInt64()
	} else {
		attributionWindow = nil
	}
	var credentials *shared.SourceTiktokMarketingAuthenticationMethod
	if r.Configuration.Credentials != nil {
		var sourceTiktokMarketingAuthenticationMethodOAuth20 *shared.SourceTiktokMarketingAuthenticationMethodOAuth20
		if r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AccessToken.ValueString()
			advertiserID := new(string)
			if !r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID.IsUnknown() && !r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID.IsNull() {
				*advertiserID = r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AdvertiserID.ValueString()
			} else {
				advertiserID = nil
			}
			appID := r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AppID.ValueString()
			authType := new(shared.SourceTiktokMarketingAuthenticationMethodOAuth20AuthType)
			if !r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType.IsNull() {
				*authType = shared.SourceTiktokMarketingAuthenticationMethodOAuth20AuthType(r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.AuthType.ValueString())
			} else {
				authType = nil
			}
			secret := r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodOAuth20.Secret.ValueString()
			sourceTiktokMarketingAuthenticationMethodOAuth20 = &shared.SourceTiktokMarketingAuthenticationMethodOAuth20{
				AccessToken:  accessToken,
				AdvertiserID: advertiserID,
				AppID:        appID,
				AuthType:     authType,
				Secret:       secret,
			}
		}
		if sourceTiktokMarketingAuthenticationMethodOAuth20 != nil {
			credentials = &shared.SourceTiktokMarketingAuthenticationMethod{
				SourceTiktokMarketingAuthenticationMethodOAuth20: sourceTiktokMarketingAuthenticationMethodOAuth20,
			}
		}
		var sourceTiktokMarketingAuthenticationMethodSandboxAccessToken *shared.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken
		if r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken != nil {
			accessToken1 := r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AccessToken.ValueString()
			advertiserId1 := r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AdvertiserID.ValueString()
			authType1 := new(shared.SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType)
			if !r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType.IsNull() {
				*authType1 = shared.SourceTiktokMarketingAuthenticationMethodSandboxAccessTokenAuthType(r.Configuration.Credentials.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			sourceTiktokMarketingAuthenticationMethodSandboxAccessToken = &shared.SourceTiktokMarketingAuthenticationMethodSandboxAccessToken{
				AccessToken:  accessToken1,
				AdvertiserID: advertiserId1,
				AuthType:     authType1,
			}
		}
		if sourceTiktokMarketingAuthenticationMethodSandboxAccessToken != nil {
			credentials = &shared.SourceTiktokMarketingAuthenticationMethod{
				SourceTiktokMarketingAuthenticationMethodSandboxAccessToken: sourceTiktokMarketingAuthenticationMethodSandboxAccessToken,
			}
		}
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	sourceType := new(shared.SourceTiktokMarketingTiktokMarketing)
	if !r.Configuration.SourceType.IsUnknown() && !r.Configuration.SourceType.IsNull() {
		*sourceType = shared.SourceTiktokMarketingTiktokMarketing(r.Configuration.SourceType.ValueString())
	} else {
		sourceType = nil
	}
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceTiktokMarketing{
		AttributionWindow: attributionWindow,
		Credentials:       credentials,
		EndDate:           endDate,
		IncludeDeleted:    includeDeleted,
		SourceType:        sourceType,
		StartDate:         startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTiktokMarketingCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTiktokMarketingResourceModel) ToGetSDKType() *shared.SourceTiktokMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTiktokMarketingResourceModel) ToUpdateSDKType() *shared.SourceTiktokMarketingPutRequest {
	attributionWindow := new(int64)
	if !r.Configuration.AttributionWindow.IsUnknown() && !r.Configuration.AttributionWindow.IsNull() {
		*attributionWindow = r.Configuration.AttributionWindow.ValueInt64()
	} else {
		attributionWindow = nil
	}
	var credentials *shared.SourceTiktokMarketingUpdateAuthenticationMethod
	if r.Configuration.Credentials != nil {
		var sourceTiktokMarketingUpdateAuthenticationMethodOAuth20 *shared.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20
		if r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20 != nil {
			accessToken := r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AccessToken.ValueString()
			advertiserID := new(string)
			if !r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AdvertiserID.IsUnknown() && !r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AdvertiserID.IsNull() {
				*advertiserID = r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AdvertiserID.ValueString()
			} else {
				advertiserID = nil
			}
			appID := r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AppID.ValueString()
			authType := new(shared.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20AuthType)
			if !r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AuthType.IsNull() {
				*authType = shared.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20AuthType(r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.AuthType.ValueString())
			} else {
				authType = nil
			}
			secret := r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20.Secret.ValueString()
			sourceTiktokMarketingUpdateAuthenticationMethodOAuth20 = &shared.SourceTiktokMarketingUpdateAuthenticationMethodOAuth20{
				AccessToken:  accessToken,
				AdvertiserID: advertiserID,
				AppID:        appID,
				AuthType:     authType,
				Secret:       secret,
			}
		}
		if sourceTiktokMarketingUpdateAuthenticationMethodOAuth20 != nil {
			credentials = &shared.SourceTiktokMarketingUpdateAuthenticationMethod{
				SourceTiktokMarketingUpdateAuthenticationMethodOAuth20: sourceTiktokMarketingUpdateAuthenticationMethodOAuth20,
			}
		}
		var sourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken *shared.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken
		if r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken != nil {
			accessToken1 := r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken.AccessToken.ValueString()
			advertiserId1 := r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken.AdvertiserID.ValueString()
			authType1 := new(shared.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessTokenAuthType)
			if !r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken.AuthType.IsUnknown() && !r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken.AuthType.IsNull() {
				*authType1 = shared.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessTokenAuthType(r.Configuration.Credentials.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken.AuthType.ValueString())
			} else {
				authType1 = nil
			}
			sourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken = &shared.SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken{
				AccessToken:  accessToken1,
				AdvertiserID: advertiserId1,
				AuthType:     authType1,
			}
		}
		if sourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken != nil {
			credentials = &shared.SourceTiktokMarketingUpdateAuthenticationMethod{
				SourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken: sourceTiktokMarketingUpdateAuthenticationMethodSandboxAccessToken,
			}
		}
	}
	endDate := new(customTypes.Date)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		endDate = customTypes.MustNewDateFromString(r.Configuration.EndDate.ValueString())
	} else {
		endDate = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	startDate := new(customTypes.Date)
	if !r.Configuration.StartDate.IsUnknown() && !r.Configuration.StartDate.IsNull() {
		startDate = customTypes.MustNewDateFromString(r.Configuration.StartDate.ValueString())
	} else {
		startDate = nil
	}
	configuration := shared.SourceTiktokMarketingUpdate{
		AttributionWindow: attributionWindow,
		Credentials:       credentials,
		EndDate:           endDate,
		IncludeDeleted:    includeDeleted,
		StartDate:         startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceTiktokMarketingPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceTiktokMarketingResourceModel) ToDeleteSDKType() *shared.SourceTiktokMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceTiktokMarketingResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceTiktokMarketingResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
