// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceTheGuardianAPIRequest struct {
	SourceTheGuardianAPIPutRequest *shared.SourceTheGuardianAPIPutRequest `request:"mediaType=application/json"`
	SourceID                       string                                 `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceTheGuardianAPIResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
