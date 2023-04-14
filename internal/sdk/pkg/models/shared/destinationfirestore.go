// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type DestinationFirestoreFirestoreEnum string

const (
	DestinationFirestoreFirestoreEnumFirestore DestinationFirestoreFirestoreEnum = "firestore"
)

func (e *DestinationFirestoreFirestoreEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "firestore":
		*e = DestinationFirestoreFirestoreEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for DestinationFirestoreFirestoreEnum: %s", s)
	}
}

// DestinationFirestore - The values required to configure the destination.
type DestinationFirestore struct {
	// The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.io/integrations/destinations/firestore">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.
	CredentialsJSON *string                           `json:"credentials_json,omitempty"`
	DestinationType DestinationFirestoreFirestoreEnum `json:"destinationType"`
	// The GCP project ID for the project containing the target BigQuery dataset.
	ProjectID string `json:"project_id"`
}
