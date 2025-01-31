// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SourceSendgridUpdate struct {
	// API Key, use <a href="https://app.sendgrid.com/settings/api_keys/">admin</a> to generate this key.
	Apikey string `json:"apikey"`
	// Start time in ISO8601 format. Any data before this time point will not be replicated.
	StartTime *time.Time `json:"start_time,omitempty"`
}
