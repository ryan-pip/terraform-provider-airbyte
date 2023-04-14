// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceMailjetSmsMailjetSmsEnum string

const (
	SourceMailjetSmsMailjetSmsEnumMailjetSms SourceMailjetSmsMailjetSmsEnum = "mailjet-sms"
)

func (e *SourceMailjetSmsMailjetSmsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "mailjet-sms":
		*e = SourceMailjetSmsMailjetSmsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMailjetSmsMailjetSmsEnum: %s", s)
	}
}

// SourceMailjetSms - The values required to configure the source.
type SourceMailjetSms struct {
	// Retrieve SMS messages created before the specified timestamp. Required format - Unix timestamp.
	EndDate    *int64                         `json:"end_date,omitempty"`
	SourceType SourceMailjetSmsMailjetSmsEnum `json:"sourceType"`
	// Retrieve SMS messages created after the specified timestamp. Required format - Unix timestamp.
	StartDate *int64 `json:"start_date,omitempty"`
	// Your access token. See <a href="https://dev.mailjet.com/sms/reference/overview/authentication">here</a>.
	Token string `json:"token"`
}
