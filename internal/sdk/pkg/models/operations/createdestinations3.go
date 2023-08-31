// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateDestinationS3Response struct {
	ContentType string
	// Successful operation
	DestinationResponse *shared.DestinationResponse
	StatusCode          int
	RawResponse         *http.Response
}
