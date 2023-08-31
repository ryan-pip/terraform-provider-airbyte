// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceLinnworksRequest struct {
	SourceLinnworksPutRequest *shared.SourceLinnworksPutRequest `request:"mediaType=application/json"`
	SourceID                  string                            `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceLinnworksResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
