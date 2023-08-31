// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy - Select value from list to pivot by
type SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy string

const (
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByCompany                     SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "COMPANY"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByAccount                     SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "ACCOUNT"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByShare                       SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "SHARE"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByCampaign                    SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CAMPAIGN"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByCreative                    SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CREATIVE"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByCampaignGroup               SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CAMPAIGN_GROUP"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByConversion                  SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CONVERSION"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByConversationNode            SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CONVERSATION_NODE"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByConversationNodeOptionIndex SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CONVERSATION_NODE_OPTION_INDEX"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByServingLocation             SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "SERVING_LOCATION"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByCardIndex                   SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "CARD_INDEX"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberCompanySize           SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_COMPANY_SIZE"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberIndustry              SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_INDUSTRY"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberSeniority             SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_SENIORITY"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberJobTitle              SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_JOB_TITLE "
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberJobFunction           SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_JOB_FUNCTION "
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberCountryV2             SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_COUNTRY_V2 "
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberRegionV2              SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_REGION_V2"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByMemberCompany               SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "MEMBER_COMPANY"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByPlacementName               SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "PLACEMENT_NAME"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotByImpressionDeviceType        SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy = "IMPRESSION_DEVICE_TYPE"
)

func (e SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy) ToPointer() *SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy {
	return &e
}

func (e *SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMPANY":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "SHARE":
		fallthrough
	case "CAMPAIGN":
		fallthrough
	case "CREATIVE":
		fallthrough
	case "CAMPAIGN_GROUP":
		fallthrough
	case "CONVERSION":
		fallthrough
	case "CONVERSATION_NODE":
		fallthrough
	case "CONVERSATION_NODE_OPTION_INDEX":
		fallthrough
	case "SERVING_LOCATION":
		fallthrough
	case "CARD_INDEX":
		fallthrough
	case "MEMBER_COMPANY_SIZE":
		fallthrough
	case "MEMBER_INDUSTRY":
		fallthrough
	case "MEMBER_SENIORITY":
		fallthrough
	case "MEMBER_JOB_TITLE ":
		fallthrough
	case "MEMBER_JOB_FUNCTION ":
		fallthrough
	case "MEMBER_COUNTRY_V2 ":
		fallthrough
	case "MEMBER_REGION_V2":
		fallthrough
	case "MEMBER_COMPANY":
		fallthrough
	case "PLACEMENT_NAME":
		fallthrough
	case "IMPRESSION_DEVICE_TYPE":
		*e = SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy: %v", v)
	}
}

// SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity - Set time granularity for report:
// ALL - Results grouped into a single result across the entire time range of the report.
// DAILY - Results grouped by day.
// MONTHLY - Results grouped by month.
// YEARLY - Results grouped by year.
type SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity string

const (
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularityAll     SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity = "ALL"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularityDaily   SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity = "DAILY"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularityMonthly SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity = "MONTHLY"
	SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularityYearly  SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity = "YEARLY"
)

func (e SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity) ToPointer() *SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity {
	return &e
}

func (e *SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "DAILY":
		fallthrough
	case "MONTHLY":
		fallthrough
	case "YEARLY":
		*e = SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity: %v", v)
	}
}

// SourceLinkedinAdsUpdateAdAnalyticsReportConfiguration - Config for custom ad Analytics Report
type SourceLinkedinAdsUpdateAdAnalyticsReportConfiguration struct {
	// The name for the report
	Name string `json:"name"`
	// Select value from list to pivot by
	PivotBy SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationPivotBy `json:"pivot_by"`
	// Set time granularity for report:
	// ALL - Results grouped into a single result across the entire time range of the report.
	// DAILY - Results grouped by day.
	// MONTHLY - Results grouped by month.
	// YEARLY - Results grouped by year.
	TimeGranularity SourceLinkedinAdsUpdateAdAnalyticsReportConfigurationTimeGranularity `json:"time_granularity"`
}

type SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod string

const (
	SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethodAccessToken SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod = "access_token"
)

func (e SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod) ToPointer() *SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod {
	return &e
}

func (e *SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod: %v", v)
	}
}

type SourceLinkedinAdsUpdateAuthenticationAccessToken struct {
	// The token value generated using the authentication code. See the <a href="https://docs.airbyte.com/integrations/sources/linkedin-ads#authentication">docs</a> to obtain yours.
	AccessToken string                                                      `json:"access_token"`
	AuthMethod  *SourceLinkedinAdsUpdateAuthenticationAccessTokenAuthMethod `json:"auth_method,omitempty"`
}

type SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod string

const (
	SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethodOAuth20 SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod = "oAuth2.0"
)

func (e SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod) ToPointer() *SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod {
	return &e
}

func (e *SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod: %v", v)
	}
}

type SourceLinkedinAdsUpdateAuthenticationOAuth20 struct {
	AuthMethod *SourceLinkedinAdsUpdateAuthenticationOAuth20AuthMethod `json:"auth_method,omitempty"`
	// The client ID of the LinkedIn Ads developer application.
	ClientID string `json:"client_id"`
	// The client secret the LinkedIn Ads developer application.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token.
	RefreshToken string `json:"refresh_token"`
}

type SourceLinkedinAdsUpdateAuthenticationType string

const (
	SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateAuthenticationOAuth20     SourceLinkedinAdsUpdateAuthenticationType = "source-linkedin-ads-update_Authentication_OAuth2.0"
	SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateAuthenticationAccessToken SourceLinkedinAdsUpdateAuthenticationType = "source-linkedin-ads-update_Authentication_Access token"
)

type SourceLinkedinAdsUpdateAuthentication struct {
	SourceLinkedinAdsUpdateAuthenticationOAuth20     *SourceLinkedinAdsUpdateAuthenticationOAuth20
	SourceLinkedinAdsUpdateAuthenticationAccessToken *SourceLinkedinAdsUpdateAuthenticationAccessToken

	Type SourceLinkedinAdsUpdateAuthenticationType
}

func CreateSourceLinkedinAdsUpdateAuthenticationSourceLinkedinAdsUpdateAuthenticationOAuth20(sourceLinkedinAdsUpdateAuthenticationOAuth20 SourceLinkedinAdsUpdateAuthenticationOAuth20) SourceLinkedinAdsUpdateAuthentication {
	typ := SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateAuthenticationOAuth20

	return SourceLinkedinAdsUpdateAuthentication{
		SourceLinkedinAdsUpdateAuthenticationOAuth20: &sourceLinkedinAdsUpdateAuthenticationOAuth20,
		Type: typ,
	}
}

func CreateSourceLinkedinAdsUpdateAuthenticationSourceLinkedinAdsUpdateAuthenticationAccessToken(sourceLinkedinAdsUpdateAuthenticationAccessToken SourceLinkedinAdsUpdateAuthenticationAccessToken) SourceLinkedinAdsUpdateAuthentication {
	typ := SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateAuthenticationAccessToken

	return SourceLinkedinAdsUpdateAuthentication{
		SourceLinkedinAdsUpdateAuthenticationAccessToken: &sourceLinkedinAdsUpdateAuthenticationAccessToken,
		Type: typ,
	}
}

func (u *SourceLinkedinAdsUpdateAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceLinkedinAdsUpdateAuthenticationOAuth20 := new(SourceLinkedinAdsUpdateAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinAdsUpdateAuthenticationOAuth20); err == nil {
		u.SourceLinkedinAdsUpdateAuthenticationOAuth20 = sourceLinkedinAdsUpdateAuthenticationOAuth20
		u.Type = SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateAuthenticationOAuth20
		return nil
	}

	sourceLinkedinAdsUpdateAuthenticationAccessToken := new(SourceLinkedinAdsUpdateAuthenticationAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinAdsUpdateAuthenticationAccessToken); err == nil {
		u.SourceLinkedinAdsUpdateAuthenticationAccessToken = sourceLinkedinAdsUpdateAuthenticationAccessToken
		u.Type = SourceLinkedinAdsUpdateAuthenticationTypeSourceLinkedinAdsUpdateAuthenticationAccessToken
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLinkedinAdsUpdateAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinAdsUpdateAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceLinkedinAdsUpdateAuthenticationOAuth20)
	}

	if u.SourceLinkedinAdsUpdateAuthenticationAccessToken != nil {
		return json.Marshal(u.SourceLinkedinAdsUpdateAuthenticationAccessToken)
	}

	return nil, nil
}

type SourceLinkedinAdsUpdate struct {
	// Specify the account IDs separated by a space, to pull the data from. Leave empty, if you want to pull the data from all associated accounts. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn Ads docs</a> for more info.
	AccountIds         []int64                                                 `json:"account_ids,omitempty"`
	AdAnalyticsReports []SourceLinkedinAdsUpdateAdAnalyticsReportConfiguration `json:"ad_analytics_reports,omitempty"`
	Credentials        *SourceLinkedinAdsUpdateAuthentication                  `json:"credentials,omitempty"`
	// UTC date in the format 2020-09-17. Any data before this date will not be replicated.
	StartDate types.Date `json:"start_date"`
}
