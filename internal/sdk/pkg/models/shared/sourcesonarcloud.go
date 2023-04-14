// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"openapi/internal/sdk/pkg/types"
)

type SourceSonarCloudSonarCloudEnum string

const (
	SourceSonarCloudSonarCloudEnumSonarCloud SourceSonarCloudSonarCloudEnum = "sonar-cloud"
)

func (e *SourceSonarCloudSonarCloudEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "sonar-cloud":
		*e = SourceSonarCloudSonarCloudEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceSonarCloudSonarCloudEnum: %s", s)
	}
}

// SourceSonarCloud - The values required to configure the source.
type SourceSonarCloud struct {
	// Comma-separated list of component keys.
	ComponentKeys []interface{} `json:"component_keys"`
	// To retrieve issues created before the given date (inclusive).
	EndDate *types.Date `json:"end_date,omitempty"`
	// Organization key. See <a href="https://docs.sonarcloud.io/appendices/project-information/#project-and-organization-keys">here</a>.
	Organization string                         `json:"organization"`
	SourceType   SourceSonarCloudSonarCloudEnum `json:"sourceType"`
	// To retrieve issues created after the given date (inclusive).
	StartDate *types.Date `json:"start_date,omitempty"`
	// Your User Token. See <a href="https://docs.sonarcloud.io/advanced-setup/user-accounts/">here</a>. The token is case sensitive.
	UserToken string `json:"user_token"`
}
