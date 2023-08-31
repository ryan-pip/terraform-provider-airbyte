// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

type SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle string

const (
	SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitleProjectSecret SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle = "Project Secret"
)

func (e SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle) ToPointer() *SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle {
	return &e
}

func (e *SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Project Secret":
		*e = SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle: %v", v)
	}
}

// SourceMixpanelUpdateAuthenticationWildcardProjectSecret - Choose how to authenticate to Mixpanel
type SourceMixpanelUpdateAuthenticationWildcardProjectSecret struct {
	// Mixpanel project secret. See the <a href="https://developer.mixpanel.com/reference/project-secret#managing-a-projects-secret">docs</a> for more information on how to obtain this.
	APISecret   string                                                              `json:"api_secret"`
	OptionTitle *SourceMixpanelUpdateAuthenticationWildcardProjectSecretOptionTitle `json:"option_title,omitempty"`
}

type SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle string

const (
	SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitleServiceAccount SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle = "Service Account"
)

func (e SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle) ToPointer() *SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle {
	return &e
}

func (e *SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "Service Account":
		*e = SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle: %v", v)
	}
}

// SourceMixpanelUpdateAuthenticationWildcardServiceAccount - Choose how to authenticate to Mixpanel
type SourceMixpanelUpdateAuthenticationWildcardServiceAccount struct {
	OptionTitle *SourceMixpanelUpdateAuthenticationWildcardServiceAccountOptionTitle `json:"option_title,omitempty"`
	// Mixpanel Service Account Secret. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.
	Secret string `json:"secret"`
	// Mixpanel Service Account Username. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.
	Username string `json:"username"`
}

type SourceMixpanelUpdateAuthenticationWildcardType string

const (
	SourceMixpanelUpdateAuthenticationWildcardTypeSourceMixpanelUpdateAuthenticationWildcardServiceAccount SourceMixpanelUpdateAuthenticationWildcardType = "source-mixpanel-update_Authentication *_Service Account"
	SourceMixpanelUpdateAuthenticationWildcardTypeSourceMixpanelUpdateAuthenticationWildcardProjectSecret  SourceMixpanelUpdateAuthenticationWildcardType = "source-mixpanel-update_Authentication *_Project Secret"
)

type SourceMixpanelUpdateAuthenticationWildcard struct {
	SourceMixpanelUpdateAuthenticationWildcardServiceAccount *SourceMixpanelUpdateAuthenticationWildcardServiceAccount
	SourceMixpanelUpdateAuthenticationWildcardProjectSecret  *SourceMixpanelUpdateAuthenticationWildcardProjectSecret

	Type SourceMixpanelUpdateAuthenticationWildcardType
}

func CreateSourceMixpanelUpdateAuthenticationWildcardSourceMixpanelUpdateAuthenticationWildcardServiceAccount(sourceMixpanelUpdateAuthenticationWildcardServiceAccount SourceMixpanelUpdateAuthenticationWildcardServiceAccount) SourceMixpanelUpdateAuthenticationWildcard {
	typ := SourceMixpanelUpdateAuthenticationWildcardTypeSourceMixpanelUpdateAuthenticationWildcardServiceAccount

	return SourceMixpanelUpdateAuthenticationWildcard{
		SourceMixpanelUpdateAuthenticationWildcardServiceAccount: &sourceMixpanelUpdateAuthenticationWildcardServiceAccount,
		Type: typ,
	}
}

func CreateSourceMixpanelUpdateAuthenticationWildcardSourceMixpanelUpdateAuthenticationWildcardProjectSecret(sourceMixpanelUpdateAuthenticationWildcardProjectSecret SourceMixpanelUpdateAuthenticationWildcardProjectSecret) SourceMixpanelUpdateAuthenticationWildcard {
	typ := SourceMixpanelUpdateAuthenticationWildcardTypeSourceMixpanelUpdateAuthenticationWildcardProjectSecret

	return SourceMixpanelUpdateAuthenticationWildcard{
		SourceMixpanelUpdateAuthenticationWildcardProjectSecret: &sourceMixpanelUpdateAuthenticationWildcardProjectSecret,
		Type: typ,
	}
}

func (u *SourceMixpanelUpdateAuthenticationWildcard) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceMixpanelUpdateAuthenticationWildcardServiceAccount := new(SourceMixpanelUpdateAuthenticationWildcardServiceAccount)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMixpanelUpdateAuthenticationWildcardServiceAccount); err == nil {
		u.SourceMixpanelUpdateAuthenticationWildcardServiceAccount = sourceMixpanelUpdateAuthenticationWildcardServiceAccount
		u.Type = SourceMixpanelUpdateAuthenticationWildcardTypeSourceMixpanelUpdateAuthenticationWildcardServiceAccount
		return nil
	}

	sourceMixpanelUpdateAuthenticationWildcardProjectSecret := new(SourceMixpanelUpdateAuthenticationWildcardProjectSecret)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceMixpanelUpdateAuthenticationWildcardProjectSecret); err == nil {
		u.SourceMixpanelUpdateAuthenticationWildcardProjectSecret = sourceMixpanelUpdateAuthenticationWildcardProjectSecret
		u.Type = SourceMixpanelUpdateAuthenticationWildcardTypeSourceMixpanelUpdateAuthenticationWildcardProjectSecret
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMixpanelUpdateAuthenticationWildcard) MarshalJSON() ([]byte, error) {
	if u.SourceMixpanelUpdateAuthenticationWildcardServiceAccount != nil {
		return json.Marshal(u.SourceMixpanelUpdateAuthenticationWildcardServiceAccount)
	}

	if u.SourceMixpanelUpdateAuthenticationWildcardProjectSecret != nil {
		return json.Marshal(u.SourceMixpanelUpdateAuthenticationWildcardProjectSecret)
	}

	return nil, nil
}

// SourceMixpanelUpdateRegion - The region of mixpanel domain instance either US or EU.
type SourceMixpanelUpdateRegion string

const (
	SourceMixpanelUpdateRegionUs SourceMixpanelUpdateRegion = "US"
	SourceMixpanelUpdateRegionEu SourceMixpanelUpdateRegion = "EU"
)

func (e SourceMixpanelUpdateRegion) ToPointer() *SourceMixpanelUpdateRegion {
	return &e
}

func (e *SourceMixpanelUpdateRegion) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "US":
		fallthrough
	case "EU":
		*e = SourceMixpanelUpdateRegion(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMixpanelUpdateRegion: %v", v)
	}
}

type SourceMixpanelUpdate struct {
	//  A period of time for attributing results to ads and the lookback period after those actions occur during which ad results are counted. Default attribution window is 5 days.
	AttributionWindow *int64 `json:"attribution_window,omitempty"`
	// Choose how to authenticate to Mixpanel
	Credentials *SourceMixpanelUpdateAuthenticationWildcard `json:"credentials,omitempty"`
	// Defines window size in days, that used to slice through data. You can reduce it, if amount of data in each window is too big for your environment.
	DateWindowSize *int64 `json:"date_window_size,omitempty"`
	// The date in the format YYYY-MM-DD. Any data after this date will not be replicated. Left empty to always sync to most recent date
	EndDate *types.Date `json:"end_date,omitempty"`
	// Your project ID number. See the <a href="https://help.mixpanel.com/hc/en-us/articles/115004490503-Project-Settings#project-id">docs</a> for more information on how to obtain this.
	ProjectID *int64 `json:"project_id,omitempty"`
	// Time zone in which integer date times are stored. The project timezone may be found in the project settings in the <a href="https://help.mixpanel.com/hc/en-us/articles/115004547203-Manage-Timezones-for-Projects-in-Mixpanel">Mixpanel console</a>.
	ProjectTimezone *string `json:"project_timezone,omitempty"`
	// The region of mixpanel domain instance either US or EU.
	Region *SourceMixpanelUpdateRegion `json:"region,omitempty"`
	// Setting this config parameter to TRUE ensures that new properties on events and engage records are captured. Otherwise new properties will be ignored.
	SelectPropertiesByDefault *bool `json:"select_properties_by_default,omitempty"`
	// The date in the format YYYY-MM-DD. Any data before this date will not be replicated. If this option is not set, the connector will replicate data from up to one year ago by default.
	StartDate *types.Date `json:"start_date,omitempty"`
}
