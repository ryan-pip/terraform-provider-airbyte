// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceRssRss string

const (
	SourceRssRssRss SourceRssRss = "rss"
)

func (e SourceRssRss) ToPointer() *SourceRssRss {
	return &e
}

func (e *SourceRssRss) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rss":
		*e = SourceRssRss(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceRssRss: %v", v)
	}
}

type SourceRss struct {
	SourceType SourceRssRss `json:"sourceType"`
	// RSS Feed URL
	URL string `json:"url"`
}
