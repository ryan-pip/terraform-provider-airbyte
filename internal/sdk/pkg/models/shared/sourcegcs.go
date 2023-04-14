// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGcsGcsEnum string

const (
	SourceGcsGcsEnumGcs SourceGcsGcsEnum = "gcs"
)

func (e *SourceGcsGcsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "gcs":
		*e = SourceGcsGcsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGcsGcsEnum: %s", s)
	}
}

// SourceGcs - The values required to configure the source.
type SourceGcs struct {
	// GCS bucket name
	GcsBucket string `json:"gcs_bucket"`
	// GCS path to data
	GcsPath string `json:"gcs_path"`
	// Enter your Google Cloud <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">service account key</a> in JSON format
	ServiceAccount string           `json:"service_account"`
	SourceType     SourceGcsGcsEnum `json:"sourceType"`
}
