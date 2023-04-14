// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type SourceInstagramInstagramEnum string

const (
	SourceInstagramInstagramEnumInstagram SourceInstagramInstagramEnum = "instagram"
)

func (e *SourceInstagramInstagramEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "instagram":
		*e = SourceInstagramInstagramEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceInstagramInstagramEnum: %s", s)
	}
}

// SourceInstagram - The values required to configure the source.
type SourceInstagram struct {
	// The value of the access token generated with <b>instagram_basic, instagram_manage_insights, pages_show_list, pages_read_engagement, Instagram Public Content Access</b> permissions. See the <a href="https://docs.airbyte.com/integrations/sources/instagram/#step-1-set-up-instagram">docs</a> for more information
	AccessToken string                       `json:"access_token"`
	SourceType  SourceInstagramInstagramEnum `json:"sourceType"`
	// The date from which you'd like to replicate data for User Insights, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
	StartDate time.Time `json:"start_date"`
}
