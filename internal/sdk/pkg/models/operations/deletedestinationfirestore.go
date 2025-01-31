// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteDestinationFirestoreRequest struct {
	DestinationID string `pathParam:"style=simple,explode=false,name=destinationId"`
}

type DeleteDestinationFirestoreResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
