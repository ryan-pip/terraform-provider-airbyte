// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// JobsResponse - List all the Jobs by connectionId.
type JobsResponse struct {
	Data     []JobResponse `json:"data"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
}
