// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceOpenweatherResourceModel) ToCreateSDKType() *shared.SourceOpenweatherCreateRequest {
	appid := r.Configuration.Appid.ValueString()
	lang := new(shared.SourceOpenweatherLanguage)
	if !r.Configuration.Lang.IsUnknown() && !r.Configuration.Lang.IsNull() {
		*lang = shared.SourceOpenweatherLanguage(r.Configuration.Lang.ValueString())
	} else {
		lang = nil
	}
	lat := r.Configuration.Lat.ValueString()
	lon := r.Configuration.Lon.ValueString()
	sourceType := shared.SourceOpenweatherOpenweather(r.Configuration.SourceType.ValueString())
	units := new(shared.SourceOpenweatherUnits)
	if !r.Configuration.Units.IsUnknown() && !r.Configuration.Units.IsNull() {
		*units = shared.SourceOpenweatherUnits(r.Configuration.Units.ValueString())
	} else {
		units = nil
	}
	configuration := shared.SourceOpenweather{
		Appid:      appid,
		Lang:       lang,
		Lat:        lat,
		Lon:        lon,
		SourceType: sourceType,
		Units:      units,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOpenweatherCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOpenweatherResourceModel) ToGetSDKType() *shared.SourceOpenweatherCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOpenweatherResourceModel) ToUpdateSDKType() *shared.SourceOpenweatherPutRequest {
	appid := r.Configuration.Appid.ValueString()
	lang := new(shared.SourceOpenweatherUpdateLanguage)
	if !r.Configuration.Lang.IsUnknown() && !r.Configuration.Lang.IsNull() {
		*lang = shared.SourceOpenweatherUpdateLanguage(r.Configuration.Lang.ValueString())
	} else {
		lang = nil
	}
	lat := r.Configuration.Lat.ValueString()
	lon := r.Configuration.Lon.ValueString()
	units := new(shared.SourceOpenweatherUpdateUnits)
	if !r.Configuration.Units.IsUnknown() && !r.Configuration.Units.IsNull() {
		*units = shared.SourceOpenweatherUpdateUnits(r.Configuration.Units.ValueString())
	} else {
		units = nil
	}
	configuration := shared.SourceOpenweatherUpdate{
		Appid: appid,
		Lang:  lang,
		Lat:   lat,
		Lon:   lon,
		Units: units,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceOpenweatherPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceOpenweatherResourceModel) ToDeleteSDKType() *shared.SourceOpenweatherCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceOpenweatherResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceOpenweatherResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
