// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteSourceAmazonAdsRequest struct {
	SourceID string `pathParam:"style=simple,explode=false,name=sourceId"`
}

type DeleteSourceAmazonAdsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
