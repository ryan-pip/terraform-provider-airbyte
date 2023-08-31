// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationDatabricksRequest struct {
	DestinationDatabricksPutRequest *shared.DestinationDatabricksPutRequest `request:"mediaType=application/json"`
	DestinationID                   string                                  `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PutDestinationDatabricksResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
