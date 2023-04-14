// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourcePosthogPosthogEnum string

const (
	SourcePosthogPosthogEnumPosthog SourcePosthogPosthogEnum = "posthog"
)

func (e *SourcePosthogPosthogEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "posthog":
		*e = SourcePosthogPosthogEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourcePosthogPosthogEnum: %s", s)
	}
}

// SourcePosthog - The values required to configure the source.
type SourcePosthog struct {
	// API Key. See the <a href="https://docs.airbyte.com/integrations/sources/posthog">docs</a> for information on how to generate this key.
	APIKey string `json:"api_key"`
	// Base PostHog url. Defaults to PostHog Cloud (https://app.posthog.com).
	BaseURL    *string                  `json:"base_url,omitempty"`
	SourceType SourcePosthogPosthogEnum `json:"sourceType"`
	// The date from which you'd like to replicate the data. Any data before this date will not be replicated.
	StartDate time.Time `json:"start_date"`
}
