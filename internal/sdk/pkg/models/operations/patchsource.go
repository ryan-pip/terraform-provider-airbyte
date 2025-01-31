// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PatchSourceRequest struct {
	SourcePatchRequest *shared.SourcePatchRequest `request:"mediaType=application/json"`
	SourceID           string                     `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PatchSourceResponse struct {
	ContentType string
	// Update a Source
	SourceResponse *shared.SourceResponse
	StatusCode     int
	RawResponse    *http.Response
}
