// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

type SourceBingAdsUpdateAuthMethod string

const (
	SourceBingAdsUpdateAuthMethodOauth20 SourceBingAdsUpdateAuthMethod = "oauth2.0"
)

func (e SourceBingAdsUpdateAuthMethod) ToPointer() *SourceBingAdsUpdateAuthMethod {
	return &e
}

func (e *SourceBingAdsUpdateAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceBingAdsUpdateAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBingAdsUpdateAuthMethod: %v", v)
	}
}

type SourceBingAdsUpdate struct {
	AuthMethod *SourceBingAdsUpdateAuthMethod `json:"auth_method,omitempty"`
	// The Client ID of your Microsoft Advertising developer application.
	ClientID string `json:"client_id"`
	// The Client Secret of your Microsoft Advertising developer application.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Developer token associated with user. See more info <a href="https://docs.microsoft.com/en-us/advertising/guides/get-started?view=bingads-13#get-developer-token"> in the docs</a>.
	DeveloperToken string `json:"developer_token"`
	// Also known as attribution or conversion window. How far into the past to look for records (in days). If your conversion window has an hours/minutes granularity, round it up to the number of days exceeding. Used only for performance report streams in incremental mode.
	LookbackWindow *int64 `json:"lookback_window,omitempty"`
	// Refresh Token to renew the expired Access Token.
	RefreshToken string `json:"refresh_token"`
	// The start date from which to begin replicating report data. Any data generated before this date will not be replicated in reports. This is a UTC date in YYYY-MM-DD format.
	ReportsStartDate types.Date `json:"reports_start_date"`
	// The Tenant ID of your Microsoft Advertising developer application. Set this to "common" unless you know you need a different value.
	TenantID *string `json:"tenant_id,omitempty"`
}
