// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type CreateSourceQualarooResponse struct {
	ContentType string
	// Successful operation
	SourceResponse *shared.SourceResponse
	StatusCode     int
	RawResponse    *http.Response
}
