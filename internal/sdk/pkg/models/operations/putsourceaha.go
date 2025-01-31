// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceAhaRequest struct {
	SourceAhaPutRequest *shared.SourceAhaPutRequest `request:"mediaType=application/json"`
	SourceID            string                      `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceAhaResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
