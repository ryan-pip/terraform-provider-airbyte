// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ScheduleTypeWithBasicEnumEnum string

const (
	ScheduleTypeWithBasicEnumEnumManual ScheduleTypeWithBasicEnumEnum = "manual"
	ScheduleTypeWithBasicEnumEnumCron   ScheduleTypeWithBasicEnumEnum = "cron"
	ScheduleTypeWithBasicEnumEnumBasic  ScheduleTypeWithBasicEnumEnum = "basic"
)

func (e *ScheduleTypeWithBasicEnumEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "manual":
		fallthrough
	case "cron":
		fallthrough
	case "basic":
		*e = ScheduleTypeWithBasicEnumEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ScheduleTypeWithBasicEnumEnum: %s", s)
	}
}
