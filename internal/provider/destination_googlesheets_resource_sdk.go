// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationGoogleSheetsResourceModel) ToCreateSDKType() *shared.DestinationGoogleSheetsCreateRequest {
	clientID := r.Configuration.Credentials.ClientID.ValueString()
	clientSecret := r.Configuration.Credentials.ClientSecret.ValueString()
	refreshToken := r.Configuration.Credentials.RefreshToken.ValueString()
	credentials := shared.DestinationGoogleSheetsAuthenticationViaGoogleOAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}
	destinationType := shared.DestinationGoogleSheetsGoogleSheets(r.Configuration.DestinationType.ValueString())
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	configuration := shared.DestinationGoogleSheets{
		Credentials:     credentials,
		DestinationType: destinationType,
		SpreadsheetID:   spreadsheetID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationGoogleSheetsCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationGoogleSheetsResourceModel) ToGetSDKType() *shared.DestinationGoogleSheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationGoogleSheetsResourceModel) ToUpdateSDKType() *shared.DestinationGoogleSheetsPutRequest {
	clientID := r.Configuration.Credentials.ClientID.ValueString()
	clientSecret := r.Configuration.Credentials.ClientSecret.ValueString()
	refreshToken := r.Configuration.Credentials.RefreshToken.ValueString()
	credentials := shared.DestinationGoogleSheetsUpdateAuthenticationViaGoogleOAuth{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
	}
	spreadsheetID := r.Configuration.SpreadsheetID.ValueString()
	configuration := shared.DestinationGoogleSheetsUpdate{
		Credentials:   credentials,
		SpreadsheetID: spreadsheetID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationGoogleSheetsPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationGoogleSheetsResourceModel) ToDeleteSDKType() *shared.DestinationGoogleSheetsCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationGoogleSheetsResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationGoogleSheetsResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
