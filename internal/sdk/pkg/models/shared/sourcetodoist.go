// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceTodoistTodoist string

const (
	SourceTodoistTodoistTodoist SourceTodoistTodoist = "todoist"
)

func (e SourceTodoistTodoist) ToPointer() *SourceTodoistTodoist {
	return &e
}

func (e *SourceTodoistTodoist) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "todoist":
		*e = SourceTodoistTodoist(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceTodoistTodoist: %v", v)
	}
}

type SourceTodoist struct {
	SourceType SourceTodoistTodoist `json:"sourceType"`
	// Your API Token. See <a href="https://todoist.com/app/settings/integrations/">here</a>. The token is case sensitive.
	Token string `json:"token"`
}
