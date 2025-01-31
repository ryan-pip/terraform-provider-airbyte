// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceSlackRequest struct {
	SourceSlackPutRequest *shared.SourceSlackPutRequest `request:"mediaType=application/json"`
	SourceID              string                        `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceSlackResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
