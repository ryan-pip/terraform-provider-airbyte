// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetDestinationTypesenseRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type GetDestinationTypesenseResponse struct {
	ContentType string
	// Get a Destination by the id in the path.
	DestinationResponse *shared.DestinationResponse
	StatusCode          int
	RawResponse         *http.Response
}
