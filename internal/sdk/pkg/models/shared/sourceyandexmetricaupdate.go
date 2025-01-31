// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
)

type SourceYandexMetricaUpdate struct {
	// Your Yandex Metrica API access token
	AuthToken string `json:"auth_token"`
	// Counter ID
	CounterID string `json:"counter_id"`
	// Starting point for your data replication, in format of "YYYY-MM-DD". If not provided will sync till most recent date.
	EndDate *types.Date `json:"end_date,omitempty"`
	// Starting point for your data replication, in format of "YYYY-MM-DD".
	StartDate types.Date `json:"start_date"`
}
