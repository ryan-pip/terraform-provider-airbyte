// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetJobRequest struct {
	JobID int64 `pathParam:"style=simple,explode=false,name=jobId"`
}

type GetJobResponse struct {
	ContentType string
	// Get a Job by the id in the path.
	JobResponse *shared.JobResponse
	StatusCode  int
	RawResponse *http.Response
}
