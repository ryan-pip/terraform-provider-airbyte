// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/models/shared"
	"net/http"
)

type PutSourceTodoistRequest struct {
	SourceTodoistPutRequest *shared.SourceTodoistPutRequest `request:"mediaType=application/json"`
	SourceID                string                          `pathParam:"style=simple,explode=false,name=sourceId"`
}

type PutSourceTodoistResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
