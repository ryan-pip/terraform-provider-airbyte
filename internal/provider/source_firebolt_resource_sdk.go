// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *SourceFireboltResourceModel) ToCreateSDKType() *shared.SourceFireboltCreateRequest {
	account := new(string)
	if !r.Configuration.Account.IsUnknown() && !r.Configuration.Account.IsNull() {
		*account = r.Configuration.Account.ValueString()
	} else {
		account = nil
	}
	database := r.Configuration.Database.ValueString()
	engine := new(string)
	if !r.Configuration.Engine.IsUnknown() && !r.Configuration.Engine.IsNull() {
		*engine = r.Configuration.Engine.ValueString()
	} else {
		engine = nil
	}
	host := new(string)
	if !r.Configuration.Host.IsUnknown() && !r.Configuration.Host.IsNull() {
		*host = r.Configuration.Host.ValueString()
	} else {
		host = nil
	}
	password := r.Configuration.Password.ValueString()
	sourceType := shared.SourceFireboltFirebolt(r.Configuration.SourceType.ValueString())
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceFirebolt{
		Account:    account,
		Database:   database,
		Engine:     engine,
		Host:       host,
		Password:   password,
		SourceType: sourceType,
		Username:   username,
	}
	name := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFireboltCreateRequest{
		Configuration: configuration,
		Name:          name,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFireboltResourceModel) ToGetSDKType() *shared.SourceFireboltCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFireboltResourceModel) ToUpdateSDKType() *shared.SourceFireboltPutRequest {
	account := new(string)
	if !r.Configuration.Account.IsUnknown() && !r.Configuration.Account.IsNull() {
		*account = r.Configuration.Account.ValueString()
	} else {
		account = nil
	}
	database := r.Configuration.Database.ValueString()
	engine := new(string)
	if !r.Configuration.Engine.IsUnknown() && !r.Configuration.Engine.IsNull() {
		*engine = r.Configuration.Engine.ValueString()
	} else {
		engine = nil
	}
	host := new(string)
	if !r.Configuration.Host.IsUnknown() && !r.Configuration.Host.IsNull() {
		*host = r.Configuration.Host.ValueString()
	} else {
		host = nil
	}
	password := r.Configuration.Password.ValueString()
	username := r.Configuration.Username.ValueString()
	configuration := shared.SourceFireboltUpdate{
		Account:  account,
		Database: database,
		Engine:   engine,
		Host:     host,
		Password: password,
		Username: username,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFireboltPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFireboltResourceModel) ToDeleteSDKType() *shared.SourceFireboltCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFireboltResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFireboltResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}
