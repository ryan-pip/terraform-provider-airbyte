// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceSlackResourceModel) ToCreateSDKType() *shared.SourceSlackCreateRequest {
	var channelFilter []string = nil
	for _, channelFilterItem := range r.Configuration.ChannelFilter {
		channelFilter = append(channelFilter, channelFilterItem.ValueString())
	}
	var credentials *shared.SourceSlackAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceSlackAuthenticationMechanismSignInViaSlackOAuth *shared.SourceSlackAuthenticationMechanismSignInViaSlackOAuth
		if r.Configuration.Credentials.SourceSlackAuthenticationMechanismSignInViaSlackOAuth != nil {
			accessToken := r.Configuration.Credentials.SourceSlackAuthenticationMechanismSignInViaSlackOAuth.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.SourceSlackAuthenticationMechanismSignInViaSlackOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceSlackAuthenticationMechanismSignInViaSlackOAuth.ClientSecret.ValueString()
			optionTitle := shared.SourceSlackAuthenticationMechanismSignInViaSlackOAuthOptionTitle(r.Configuration.Credentials.SourceSlackAuthenticationMechanismSignInViaSlackOAuth.OptionTitle.ValueString())
			sourceSlackAuthenticationMechanismSignInViaSlackOAuth = &shared.SourceSlackAuthenticationMechanismSignInViaSlackOAuth{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				OptionTitle:  optionTitle,
			}
		}
		if sourceSlackAuthenticationMechanismSignInViaSlackOAuth != nil {
			credentials = &shared.SourceSlackAuthenticationMechanism{
				SourceSlackAuthenticationMechanismSignInViaSlackOAuth: sourceSlackAuthenticationMechanismSignInViaSlackOAuth,
			}
		}
		var sourceSlackAuthenticationMechanismAPIToken *shared.SourceSlackAuthenticationMechanismAPIToken
		if r.Configuration.Credentials.SourceSlackAuthenticationMechanismAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceSlackAuthenticationMechanismAPIToken.APIToken.ValueString()
			optionTitle1 := shared.SourceSlackAuthenticationMechanismAPITokenOptionTitle(r.Configuration.Credentials.SourceSlackAuthenticationMechanismAPIToken.OptionTitle.ValueString())
			sourceSlackAuthenticationMechanismAPIToken = &shared.SourceSlackAuthenticationMechanismAPIToken{
				APIToken:    apiToken,
				OptionTitle: optionTitle1,
			}
		}
		if sourceSlackAuthenticationMechanismAPIToken != nil {
			credentials = &shared.SourceSlackAuthenticationMechanism{
				SourceSlackAuthenticationMechanismAPIToken: sourceSlackAuthenticationMechanismAPIToken,
			}
		}
	}
	joinChannels := r.Configuration.JoinChannels.ValueBool()
	lookbackWindow := r.Configuration.LookbackWindow.ValueInt64()
	sourceType := shared.SourceSlackSlack(r.Configuration.SourceType.ValueString())
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSlack{
		ChannelFilter:  channelFilter,
		Credentials:    credentials,
		JoinChannels:   joinChannels,
		LookbackWindow: lookbackWindow,
		SourceType:     sourceType,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSlackCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSlackResourceModel) ToGetSDKType() *shared.SourceSlackCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSlackResourceModel) ToUpdateSDKType() *shared.SourceSlackPutRequest {
	var channelFilter []string = nil
	for _, channelFilterItem := range r.Configuration.ChannelFilter {
		channelFilter = append(channelFilter, channelFilterItem.ValueString())
	}
	var credentials *shared.SourceSlackUpdateAuthenticationMechanism
	if r.Configuration.Credentials != nil {
		var sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth *shared.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth
		if r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth != nil {
			accessToken := r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth.AccessToken.ValueString()
			clientID := r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth.ClientID.ValueString()
			clientSecret := r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth.ClientSecret.ValueString()
			optionTitle := shared.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuthOptionTitle(r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth.OptionTitle.ValueString())
			sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth = &shared.SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth{
				AccessToken:  accessToken,
				ClientID:     clientID,
				ClientSecret: clientSecret,
				OptionTitle:  optionTitle,
			}
		}
		if sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth != nil {
			credentials = &shared.SourceSlackUpdateAuthenticationMechanism{
				SourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth: sourceSlackUpdateAuthenticationMechanismSignInViaSlackOAuth,
			}
		}
		var sourceSlackUpdateAuthenticationMechanismAPIToken *shared.SourceSlackUpdateAuthenticationMechanismAPIToken
		if r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismAPIToken != nil {
			apiToken := r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismAPIToken.APIToken.ValueString()
			optionTitle1 := shared.SourceSlackUpdateAuthenticationMechanismAPITokenOptionTitle(r.Configuration.Credentials.SourceSlackUpdateAuthenticationMechanismAPIToken.OptionTitle.ValueString())
			sourceSlackUpdateAuthenticationMechanismAPIToken = &shared.SourceSlackUpdateAuthenticationMechanismAPIToken{
				APIToken:    apiToken,
				OptionTitle: optionTitle1,
			}
		}
		if sourceSlackUpdateAuthenticationMechanismAPIToken != nil {
			credentials = &shared.SourceSlackUpdateAuthenticationMechanism{
				SourceSlackUpdateAuthenticationMechanismAPIToken: sourceSlackUpdateAuthenticationMechanismAPIToken,
			}
		}
	}
	joinChannels := r.Configuration.JoinChannels.ValueBool()
	lookbackWindow := r.Configuration.LookbackWindow.ValueInt64()
	startDate, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceSlackUpdate{
		ChannelFilter:  channelFilter,
		Credentials:    credentials,
		JoinChannels:   joinChannels,
		LookbackWindow: lookbackWindow,
		StartDate:      startDate,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceSlackPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceSlackResourceModel) ToDeleteSDKType() *shared.SourceSlackCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceSlackResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceSlackResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
