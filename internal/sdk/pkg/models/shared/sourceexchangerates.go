// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"openapi/internal/sdk/pkg/types"
)

type SourceExchangeRatesExchangeRatesEnum string

const (
	SourceExchangeRatesExchangeRatesEnumExchangeRates SourceExchangeRatesExchangeRatesEnum = "exchange-rates"
)

func (e *SourceExchangeRatesExchangeRatesEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "exchange-rates":
		*e = SourceExchangeRatesExchangeRatesEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceExchangeRatesExchangeRatesEnum: %s", s)
	}
}

// SourceExchangeRates - The values required to configure the source.
type SourceExchangeRates struct {
	// Your API Key. See <a href="https://apilayer.com/marketplace/exchangerates_data-api">here</a>. The key is case sensitive.
	AccessKey string `json:"access_key"`
	// ISO reference currency. See <a href="https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html">here</a>. Free plan doesn't support Source Currency Switching, default base currency is EUR
	Base *string `json:"base,omitempty"`
	// Ignore weekends? (Exchanges don't run on weekends)
	IgnoreWeekends *bool                                `json:"ignore_weekends,omitempty"`
	SourceType     SourceExchangeRatesExchangeRatesEnum `json:"sourceType"`
	// Start getting data from that date.
	StartDate types.Date `json:"start_date"`
}
