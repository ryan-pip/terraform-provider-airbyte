// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
	"time"
)

type ListJobsRequest struct {
	// Filter the Jobs by connectionId.
	ConnectionID *string `queryParam:"style=form,explode=true,name=connectionId"`
	// The end date to filter by
	CreatedAtEnd *time.Time `queryParam:"style=form,explode=true,name=createdAtEnd"`
	// The start date to filter by
	CreatedAtStart *time.Time `queryParam:"style=form,explode=true,name=createdAtStart"`
	// Filter the Jobs by jobType.
	JobType *shared.JobTypeEnum `queryParam:"style=form,explode=true,name=jobType"`
	// Set the limit on the number of Jobs returned. The default is 20 Jobs.
	Limit *int `queryParam:"style=form,explode=true,name=limit"`
	// Set the offset to start at when returning Jobs. The default is 0.
	Offset *int `queryParam:"style=form,explode=true,name=offset"`
	// The field and method to use for ordering. Currently allowed are createdAt and updatedAt.
	OrderBy *string `queryParam:"style=form,explode=true,name=orderBy"`
	// The Job status you want to filter by
	Status *shared.JobStatusEnum `queryParam:"style=form,explode=true,name=status"`
	// The end date to filter by
	UpdatedAtEnd *time.Time `queryParam:"style=form,explode=true,name=updatedAtEnd"`
	// The start date to filter by
	UpdatedAtStart *time.Time `queryParam:"style=form,explode=true,name=updatedAtStart"`
	// The UUIDs of the workspaces you wish to list jobs for. Empty list will retrieve all allowed workspaces.
	WorkspaceIds []string `queryParam:"style=form,explode=true,name=workspaceIds"`
}

type ListJobsResponse struct {
	ContentType string
	// List all the Jobs by connectionId.
	JobsResponse *shared.JobsResponse
	StatusCode   int
	RawResponse  *http.Response
}
