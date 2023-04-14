// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type JobStatusEnumEnum string

const (
	JobStatusEnumEnumPending    JobStatusEnumEnum = "pending"
	JobStatusEnumEnumRunning    JobStatusEnumEnum = "running"
	JobStatusEnumEnumIncomplete JobStatusEnumEnum = "incomplete"
	JobStatusEnumEnumFailed     JobStatusEnumEnum = "failed"
	JobStatusEnumEnumSucceeded  JobStatusEnumEnum = "succeeded"
	JobStatusEnumEnumCancelled  JobStatusEnumEnum = "cancelled"
)

func (e *JobStatusEnumEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "pending":
		fallthrough
	case "running":
		fallthrough
	case "incomplete":
		fallthrough
	case "failed":
		fallthrough
	case "succeeded":
		fallthrough
	case "cancelled":
		*e = JobStatusEnumEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for JobStatusEnumEnum: %s", s)
	}
}
