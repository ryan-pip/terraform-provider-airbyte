// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ConnectionSchedule - schedule for when the the connection should run, per the schedule type
type ConnectionSchedule struct {
	CronExpression *string          `json:"cronExpression,omitempty"`
	ScheduleType   ScheduleTypeEnum `json:"scheduleType"`
}
