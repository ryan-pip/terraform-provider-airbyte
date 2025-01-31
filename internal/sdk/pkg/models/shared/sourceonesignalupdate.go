// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SourceOnesignalUpdateApplications struct {
	AppAPIKey string  `json:"app_api_key"`
	AppID     string  `json:"app_id"`
	AppName   *string `json:"app_name,omitempty"`
}

type SourceOnesignalUpdate struct {
	// Applications keys, see the <a href="https://documentation.onesignal.com/docs/accounts-and-keys">docs</a> for more information on how to obtain this data
	Applications []SourceOnesignalUpdateApplications `json:"applications"`
	// Comma-separated list of names and the value (sum/count) for the returned outcome data. See the <a href="https://documentation.onesignal.com/reference/view-outcomes">docs</a> for more details
	OutcomeNames string `json:"outcome_names"`
	// The date from which you'd like to replicate data for OneSignal API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// OneSignal User Auth Key, see the <a href="https://documentation.onesignal.com/docs/accounts-and-keys#user-auth-key">docs</a> for more information on how to obtain this key.
	UserAuthKey string `json:"user_auth_key"`
}
