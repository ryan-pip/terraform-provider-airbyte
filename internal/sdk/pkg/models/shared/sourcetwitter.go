// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceTwitterTwitterEnum string

const (
	SourceTwitterTwitterEnumTwitter SourceTwitterTwitterEnum = "twitter"
)

func (e *SourceTwitterTwitterEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "twitter":
		*e = SourceTwitterTwitterEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTwitterTwitterEnum: %s", s)
	}
}

// SourceTwitter - The values required to configure the source.
type SourceTwitter struct {
	// App only Bearer Token. See the <a href="https://developer.twitter.com/en/docs/authentication/oauth-2-0/bearer-tokens">docs</a> for more information on how to obtain this token.
	APIKey string `json:"api_key"`
	// The end date for retrieving tweets must be a minimum of 10 seconds prior to the request time.
	EndDate *time.Time `json:"end_date,omitempty"`
	// Query for matching Tweets. You can learn how to build this query by reading <a href="https://developer.twitter.com/en/docs/twitter-api/tweets/search/integrate/build-a-query"> build a query guide </a>.
	Query      string                   `json:"query"`
	SourceType SourceTwitterTwitterEnum `json:"sourceType"`
	// The start date for retrieving tweets cannot be more than 7 days in the past.
	StartDate *time.Time `json:"start_date,omitempty"`
}
