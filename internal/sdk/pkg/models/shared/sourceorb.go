// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceOrbOrb string

const (
	SourceOrbOrbOrb SourceOrbOrb = "orb"
)

func (e SourceOrbOrb) ToPointer() *SourceOrbOrb {
	return &e
}

func (e *SourceOrbOrb) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "orb":
		*e = SourceOrbOrb(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceOrbOrb: %v", v)
	}
}

type SourceOrb struct {
	// Orb API Key, issued from the Orb admin console.
	APIKey string `json:"api_key"`
	// When set to N, the connector will always refresh resources created within the past N days. By default, updated objects that are not newly created are not incrementally synced.
	LookbackWindowDays *int64 `json:"lookback_window_days,omitempty"`
	// Property key names to extract from all events, in order to enrich ledger entries corresponding to an event deduction.
	NumericEventPropertiesKeys []string `json:"numeric_event_properties_keys,omitempty"`
	// Orb Plan ID to filter subscriptions that should have usage fetched.
	PlanID     *string      `json:"plan_id,omitempty"`
	SourceType SourceOrbOrb `json:"sourceType"`
	// UTC date and time in the format 2022-03-01T00:00:00Z. Any data with created_at before this data will not be synced. For Subscription Usage, this becomes the `timeframe_start` API parameter.
	StartDate string `json:"start_date"`
	// Property key names to extract from all events, in order to enrich ledger entries corresponding to an event deduction.
	StringEventPropertiesKeys []string `json:"string_event_properties_keys,omitempty"`
	// Property key name to group subscription usage by.
	SubscriptionUsageGroupingKey *string `json:"subscription_usage_grouping_key,omitempty"`
}
