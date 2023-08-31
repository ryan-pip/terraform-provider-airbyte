// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PatchDestinationRequest struct {
	DestinationPatchRequest *shared.DestinationPatchRequest `request:"mediaType=application/json"`
	DestinationID           string                          `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PatchDestinationResponse struct {
	ContentType string
	// Update a Destination
	DestinationResponse *shared.DestinationResponse
	StatusCode          int
	RawResponse         *http.Response
}
