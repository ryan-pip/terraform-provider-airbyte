// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutDestinationFireboltRequest struct {
	DestinationFireboltPutRequest *shared.DestinationFireboltPutRequest `request:"mediaType=application/json"`
	DestinationID                 string                                `pathParam:"style=simple,explode=false,name=destinationId"`
}

type PutDestinationFireboltResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
