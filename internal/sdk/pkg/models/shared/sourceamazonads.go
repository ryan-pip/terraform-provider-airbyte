// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceAmazonAdsAuthType string

const (
	SourceAmazonAdsAuthTypeOauth20 SourceAmazonAdsAuthType = "oauth2.0"
)

func (e SourceAmazonAdsAuthType) ToPointer() *SourceAmazonAdsAuthType {
	return &e
}

func (e *SourceAmazonAdsAuthType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oauth2.0":
		*e = SourceAmazonAdsAuthType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsAuthType: %v", v)
	}
}

// SourceAmazonAdsRegion - Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.
type SourceAmazonAdsRegion string

const (
	SourceAmazonAdsRegionNa SourceAmazonAdsRegion = "NA"
	SourceAmazonAdsRegionEu SourceAmazonAdsRegion = "EU"
	SourceAmazonAdsRegionFe SourceAmazonAdsRegion = "FE"
)

func (e SourceAmazonAdsRegion) ToPointer() *SourceAmazonAdsRegion {
	return &e
}

func (e *SourceAmazonAdsRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NA":
		fallthrough
	case "EU":
		fallthrough
	case "FE":
		*e = SourceAmazonAdsRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsRegion: %v", v)
	}
}

type SourceAmazonAdsReportRecordTypes string

const (
	SourceAmazonAdsReportRecordTypesAdGroups      SourceAmazonAdsReportRecordTypes = "adGroups"
	SourceAmazonAdsReportRecordTypesAsins         SourceAmazonAdsReportRecordTypes = "asins"
	SourceAmazonAdsReportRecordTypesAsinsKeywords SourceAmazonAdsReportRecordTypes = "asins_keywords"
	SourceAmazonAdsReportRecordTypesAsinsTargets  SourceAmazonAdsReportRecordTypes = "asins_targets"
	SourceAmazonAdsReportRecordTypesCampaigns     SourceAmazonAdsReportRecordTypes = "campaigns"
	SourceAmazonAdsReportRecordTypesKeywords      SourceAmazonAdsReportRecordTypes = "keywords"
	SourceAmazonAdsReportRecordTypesProductAds    SourceAmazonAdsReportRecordTypes = "productAds"
	SourceAmazonAdsReportRecordTypesTargets       SourceAmazonAdsReportRecordTypes = "targets"
)

func (e SourceAmazonAdsReportRecordTypes) ToPointer() *SourceAmazonAdsReportRecordTypes {
	return &e
}

func (e *SourceAmazonAdsReportRecordTypes) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "adGroups":
		fallthrough
	case "asins":
		fallthrough
	case "asins_keywords":
		fallthrough
	case "asins_targets":
		fallthrough
	case "campaigns":
		fallthrough
	case "keywords":
		fallthrough
	case "productAds":
		fallthrough
	case "targets":
		*e = SourceAmazonAdsReportRecordTypes(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsReportRecordTypes: %v", v)
	}
}

type SourceAmazonAdsAmazonAds string

const (
	SourceAmazonAdsAmazonAdsAmazonAds SourceAmazonAdsAmazonAds = "amazon-ads"
)

func (e SourceAmazonAdsAmazonAds) ToPointer() *SourceAmazonAdsAmazonAds {
	return &e
}

func (e *SourceAmazonAdsAmazonAds) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "amazon-ads":
		*e = SourceAmazonAdsAmazonAds(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsAmazonAds: %v", v)
	}
}

type SourceAmazonAdsStateFilter string

const (
	SourceAmazonAdsStateFilterEnabled  SourceAmazonAdsStateFilter = "enabled"
	SourceAmazonAdsStateFilterPaused   SourceAmazonAdsStateFilter = "paused"
	SourceAmazonAdsStateFilterArchived SourceAmazonAdsStateFilter = "archived"
)

func (e SourceAmazonAdsStateFilter) ToPointer() *SourceAmazonAdsStateFilter {
	return &e
}

func (e *SourceAmazonAdsStateFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "paused":
		fallthrough
	case "archived":
		*e = SourceAmazonAdsStateFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceAmazonAdsStateFilter: %v", v)
	}
}

type SourceAmazonAds struct {
	AuthType *SourceAmazonAdsAuthType `json:"auth_type,omitempty"`
	// The client ID of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.
	ClientID string `json:"client_id"`
	// The client secret of your Amazon Ads developer application. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens#retrieve-your-client-id-and-client-secret">docs</a> for more information.
	ClientSecret string `json:"client_secret"`
	// The amount of days to go back in time to get the updated data from Amazon Ads
	LookBackWindow *int64 `json:"look_back_window,omitempty"`
	// Profile IDs you want to fetch data for. See <a href="https://advertising.amazon.com/API/docs/en-us/concepts/authorization/profiles">docs</a> for more details.
	Profiles []int64 `json:"profiles,omitempty"`
	// Amazon Ads refresh token. See the <a href="https://advertising.amazon.com/API/docs/en-us/get-started/generate-api-tokens">docs</a> for more information on how to obtain this token.
	RefreshToken string `json:"refresh_token"`
	// Region to pull data from (EU/NA/FE). See <a href="https://advertising.amazon.com/API/docs/en-us/info/api-overview#api-endpoints">docs</a> for more details.
	Region *SourceAmazonAdsRegion `json:"region,omitempty"`
	// Optional configuration which accepts an array of string of record types. Leave blank for default behaviour to pull all report types. Use this config option only if you want to pull specific report type(s). See <a href="https://advertising.amazon.com/API/docs/en-us/reporting/v2/report-types">docs</a> for more details
	ReportRecordTypes []SourceAmazonAdsReportRecordTypes `json:"report_record_types,omitempty"`
	SourceType        SourceAmazonAdsAmazonAds           `json:"sourceType"`
	// The Start date for collecting reports, should not be more than 60 days in the past. In YYYY-MM-DD format
	StartDate *string `json:"start_date,omitempty"`
	// Reflects the state of the Display, Product, and Brand Campaign streams as enabled, paused, or archived. If you do not populate this field, it will be ignored completely.
	StateFilter []SourceAmazonAdsStateFilter `json:"state_filter,omitempty"`
}
