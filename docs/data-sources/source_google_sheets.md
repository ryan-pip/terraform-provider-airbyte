---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_sheets Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGoogleSheets DataSource
---

# airbyte_source_google_sheets (Data Source)

SourceGoogleSheets DataSource

## Example Usage

```terraform
data "airbyte_source_google_sheets" "my_source_googlesheets" {
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

- `credentials` (Attributes) Credentials for connecting to the Google Sheets API (see [below for nested schema](#nestedatt--configuration--credentials))
- `names_conversion` (Boolean) Enables the conversion of column names to a standardized, SQL-compliant format. For example, 'My Name' -> 'my_name'. Enable this option if your destination is SQL-based.
- `row_batch_size` (Number) The number of rows fetched when making a Google Sheet API call. Defaults to 200.
- `source_type` (String) must be one of ["google-sheets"]
- `spreadsheet_id` (String) Enter the link to the Google spreadsheet you want to sync. To copy the link, click the 'Share' button in the top-right corner of the spreadsheet, then click 'Copy link'.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_google_sheets_authentication_authenticate_via_google_o_auth` (Attributes) Credentials for connecting to the Google Sheets API (see [below for nested schema](#nestedatt--configuration--credentials--source_google_sheets_authentication_authenticate_via_google_o_auth))
- `source_google_sheets_authentication_service_account_key_authentication` (Attributes) Credentials for connecting to the Google Sheets API (see [below for nested schema](#nestedatt--configuration--credentials--source_google_sheets_authentication_service_account_key_authentication))
- `source_google_sheets_update_authentication_authenticate_via_google_o_auth` (Attributes) Credentials for connecting to the Google Sheets API (see [below for nested schema](#nestedatt--configuration--credentials--source_google_sheets_update_authentication_authenticate_via_google_o_auth))
- `source_google_sheets_update_authentication_service_account_key_authentication` (Attributes) Credentials for connecting to the Google Sheets API (see [below for nested schema](#nestedatt--configuration--credentials--source_google_sheets_update_authentication_service_account_key_authentication))

<a id="nestedatt--configuration--credentials--source_google_sheets_authentication_authenticate_via_google_o_auth"></a>
### Nested Schema for `configuration.credentials.source_google_sheets_authentication_authenticate_via_google_o_auth`

Read-Only:

- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) Enter your Google application's Client ID. See <a href='https://developers.google.com/identity/protocols/oauth2'>Google's documentation</a> for more information.
- `client_secret` (String) Enter your Google application's Client Secret. See <a href='https://developers.google.com/identity/protocols/oauth2'>Google's documentation</a> for more information.
- `refresh_token` (String) Enter your Google application's refresh token. See <a href='https://developers.google.com/identity/protocols/oauth2'>Google's documentation</a> for more information.


<a id="nestedatt--configuration--credentials--source_google_sheets_authentication_service_account_key_authentication"></a>
### Nested Schema for `configuration.credentials.source_google_sheets_authentication_service_account_key_authentication`

Read-Only:

- `auth_type` (String) must be one of ["Service"]
- `service_account_info` (String) The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">here</a>.


<a id="nestedatt--configuration--credentials--source_google_sheets_update_authentication_authenticate_via_google_o_auth"></a>
### Nested Schema for `configuration.credentials.source_google_sheets_update_authentication_authenticate_via_google_o_auth`

Read-Only:

- `auth_type` (String) must be one of ["Client"]
- `client_id` (String) Enter your Google application's Client ID. See <a href='https://developers.google.com/identity/protocols/oauth2'>Google's documentation</a> for more information.
- `client_secret` (String) Enter your Google application's Client Secret. See <a href='https://developers.google.com/identity/protocols/oauth2'>Google's documentation</a> for more information.
- `refresh_token` (String) Enter your Google application's refresh token. See <a href='https://developers.google.com/identity/protocols/oauth2'>Google's documentation</a> for more information.


<a id="nestedatt--configuration--credentials--source_google_sheets_update_authentication_service_account_key_authentication"></a>
### Nested Schema for `configuration.credentials.source_google_sheets_update_authentication_service_account_key_authentication`

Read-Only:

- `auth_type` (String) must be one of ["Service"]
- `service_account_info` (String) The JSON key of the service account to use for authorization. Read more <a href="https://cloud.google.com/iam/docs/creating-managing-service-account-keys#creating_service_account_keys">here</a>.


