// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceBigcommerceBigcommerceEnum string

const (
	SourceBigcommerceBigcommerceEnumBigcommerce SourceBigcommerceBigcommerceEnum = "bigcommerce"
)

func (e *SourceBigcommerceBigcommerceEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "bigcommerce":
		*e = SourceBigcommerceBigcommerceEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceBigcommerceBigcommerceEnum: %s", s)
	}
}

// SourceBigcommerce - The values required to configure the source.
type SourceBigcommerce struct {
	// Access Token for making authenticated requests.
	AccessToken string                           `json:"access_token"`
	SourceType  SourceBigcommerceBigcommerceEnum `json:"sourceType"`
	// The date you would like to replicate data. Format: YYYY-MM-DD.
	StartDate string `json:"start_date"`
	// The hash code of the store. For https://api.bigcommerce.com/stores/HASH_CODE/v3/, The store's hash code is 'HASH_CODE'.
	StoreHash string `json:"store_hash"`
}
