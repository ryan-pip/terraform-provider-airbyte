// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SourceStripeUpdate struct {
	// Your Stripe account ID (starts with 'acct_', find yours <a href="https://dashboard.stripe.com/settings/account">here</a>).
	AccountID string `json:"account_id"`
	// Stripe API key (usually starts with 'sk_live_'; find yours <a href="https://dashboard.stripe.com/apikeys">here</a>).
	ClientSecret string `json:"client_secret"`
	// When set, the connector will always re-export data from the past N days, where N is the value set here. This is useful if your data is frequently updated after creation. More info <a href="https://docs.airbyte.com/integrations/sources/stripe#requirements">here</a>
	LookbackWindowDays *int64 `json:"lookback_window_days,omitempty"`
	// The time increment used by the connector when requesting data from the Stripe API. The bigger the value is, the less requests will be made and faster the sync will be. On the other hand, the more seldom the state is persisted.
	SliceRange *int64 `json:"slice_range,omitempty"`
	// UTC date and time in the format 2017-01-25T00:00:00Z. Only data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
}
