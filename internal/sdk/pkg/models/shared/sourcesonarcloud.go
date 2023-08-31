// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/ryan-pip/terraform-provider-airbyte/internal/sdk/pkg/types"
	"encoding/json"
	"fmt"
)

type SourceSonarCloudSonarCloud string

const (
	SourceSonarCloudSonarCloudSonarCloud SourceSonarCloudSonarCloud = "sonar-cloud"
)

func (e SourceSonarCloudSonarCloud) ToPointer() *SourceSonarCloudSonarCloud {
	return &e
}

func (e *SourceSonarCloudSonarCloud) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sonar-cloud":
		*e = SourceSonarCloudSonarCloud(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSonarCloudSonarCloud: %v", v)
	}
}

type SourceSonarCloud struct {
	// Comma-separated list of component keys.
	ComponentKeys []interface{} `json:"component_keys"`
	// To retrieve issues created before the given date (inclusive).
	EndDate *types.Date `json:"end_date,omitempty"`
	// Organization key. See <a href="https://docs.sonarcloud.io/appendices/project-information/#project-and-organization-keys">here</a>.
	Organization string                     `json:"organization"`
	SourceType   SourceSonarCloudSonarCloud `json:"sourceType"`
	// To retrieve issues created after the given date (inclusive).
	StartDate *types.Date `json:"start_date,omitempty"`
	// Your User Token. See <a href="https://docs.sonarcloud.io/advanced-setup/user-accounts/">here</a>. The token is case sensitive.
	UserToken string `json:"user_token"`
}
