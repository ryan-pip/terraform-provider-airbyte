---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_analytics_data_api Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGoogleAnalyticsDataAPI DataSource
---

# airbyte_source_google_analytics_data_api (Data Source)

SourceGoogleAnalyticsDataAPI DataSource

## Example Usage

```terraform
data "airbyte_source_google_analytics_data_api" "my_source_googleanalyticsdataapi" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `credentials` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials))
- `custom_reports` (String) A JSON array describing the custom reports you want to sync from Google Analytics. See <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#custom-reports">the documentation</a> for more information about the exact format you can use to fill out this field.
- `date_ranges_start_date` (String) The start date from which to replicate report data in the format YYYY-MM-DD. Data generated before this date will not be included in the report. Not applied to custom Cohort reports.
- `property_id` (String) The Property ID is a unique number assigned to each property in Google Analytics, found in your GA4 property URL. This ID allows the connector to track the specific events associated with your property. Refer to the <a href='https://developers.google.com/analytics/devguides/reporting/data/v1/property-id#what_is_my_property_id'>Google Analytics documentation</a> to locate your property ID.
- `source_type` (String) must be one of ["google-analytics-data-api"]
- `window_in_days` (Number) The interval in days for each data request made to the Google Analytics API. A larger value speeds up data sync, but increases the chance of data sampling, which may result in inaccuracies. We recommend a value of 1 to minimize sampling, unless speed is an absolute priority over accuracy. Acceptable values range from 1 to 364. Does not apply to custom Cohort reports. More information is available in <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api">the documentation</a>.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_google_analytics_data_api_credentials_authenticate_via_google_oauth` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_authenticate_via_google_oauth))
- `source_google_analytics_data_api_credentials_service_account_key_authentication` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_service_account_key_authentication))
- `source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth))
- `source_google_analytics_data_api_update_credentials_service_account_key_authentication` (Attributes) Credentials for the service (see [below for nested schema](#nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_service_account_key_authentication))

<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_authenticate_via_google_oauth"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_credentials_authenticate_via_google_oauth`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) The Client ID of your Google Analytics developer application.
- `client_secret` (String) The Client Secret of your Google Analytics developer application.
- `refresh_token` (String) The token for obtaining a new access token.


<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_credentials_service_account_key_authentication"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_credentials_service_account_key_authentication`

Read-Only:

- `auth_type` (String) must be one of ["Service"]
- `credentials_json` (String) The JSON key linked to the service account used for authorization. For steps on obtaining this key, refer to <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#setup-guide">the setup guide</a>.


<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_update_credentials_authenticate_via_google_oauth`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests.
- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) The Client ID of your Google Analytics developer application.
- `client_secret` (String) The Client Secret of your Google Analytics developer application.
- `refresh_token` (String) The token for obtaining a new access token.


<a id="nestedatt--configuration--credentials--source_google_analytics_data_api_update_credentials_service_account_key_authentication"></a>
### Nested Schema for `configuration.credentials.source_google_analytics_data_api_update_credentials_service_account_key_authentication`

Read-Only:

- `auth_type` (String) must be one of ["Service"]
- `credentials_json` (String) The JSON key linked to the service account used for authorization. For steps on obtaining this key, refer to <a href="https://docs.airbyte.com/integrations/sources/google-analytics-data-api/#setup-guide">the setup guide</a>.


