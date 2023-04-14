// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/internal/sdk/pkg/models/shared"
)

type ListWorkspacesRequest struct {
	// Include deleted workspaces in the returned results.
	IncludeDeleted *bool `queryParam:"style=form,explode=true,name=includeDeleted"`
	// Set the limit on the number of workspaces returned. The default is 20.
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	// Set the offset to start at when returning workspaces. The default is 0
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// The UUIDs of the workspaces you wish to fetch. Empty list will retrieve all allowed workspaces.
	WorkspaceIds []string `queryParam:"style=form,explode=true,name=workspaceIds"`
}

type ListWorkspacesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Successful operation
	WorkspacesResponse *shared.WorkspacesResponse
}
