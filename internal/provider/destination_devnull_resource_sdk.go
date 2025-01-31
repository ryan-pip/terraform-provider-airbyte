// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationDevNullResourceModel) ToCreateSDKType() *shared.DestinationDevNullCreateRequest {
	destinationType := shared.DestinationDevNullDevNull(r.Configuration.DestinationType.ValueString())
	var testDestination shared.DestinationDevNullTestDestination
	var destinationDevNullTestDestinationSilent *shared.DestinationDevNullTestDestinationSilent
	if r.Configuration.TestDestination.DestinationDevNullTestDestinationSilent != nil {
		testDestinationType := shared.DestinationDevNullTestDestinationSilentTestDestinationType(r.Configuration.TestDestination.DestinationDevNullTestDestinationSilent.TestDestinationType.ValueString())
		destinationDevNullTestDestinationSilent = &shared.DestinationDevNullTestDestinationSilent{
			TestDestinationType: testDestinationType,
		}
	}
	if destinationDevNullTestDestinationSilent != nil {
		testDestination = shared.DestinationDevNullTestDestination{
			DestinationDevNullTestDestinationSilent: destinationDevNullTestDestinationSilent,
		}
	}
	configuration := shared.DestinationDevNull{
		DestinationType: destinationType,
		TestDestination: testDestination,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDevNullCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDevNullResourceModel) ToGetSDKType() *shared.DestinationDevNullCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDevNullResourceModel) ToUpdateSDKType() *shared.DestinationDevNullPutRequest {
	var testDestination shared.DestinationDevNullUpdateTestDestination
	var destinationDevNullUpdateTestDestinationSilent *shared.DestinationDevNullUpdateTestDestinationSilent
	if r.Configuration.TestDestination.DestinationDevNullUpdateTestDestinationSilent != nil {
		testDestinationType := shared.DestinationDevNullUpdateTestDestinationSilentTestDestinationType(r.Configuration.TestDestination.DestinationDevNullUpdateTestDestinationSilent.TestDestinationType.ValueString())
		destinationDevNullUpdateTestDestinationSilent = &shared.DestinationDevNullUpdateTestDestinationSilent{
			TestDestinationType: testDestinationType,
		}
	}
	if destinationDevNullUpdateTestDestinationSilent != nil {
		testDestination = shared.DestinationDevNullUpdateTestDestination{
			DestinationDevNullUpdateTestDestinationSilent: destinationDevNullUpdateTestDestinationSilent,
		}
	}
	configuration := shared.DestinationDevNullUpdate{
		TestDestination: testDestination,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationDevNullPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationDevNullResourceModel) ToDeleteSDKType() *shared.DestinationDevNullCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationDevNullResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationDevNullResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}
