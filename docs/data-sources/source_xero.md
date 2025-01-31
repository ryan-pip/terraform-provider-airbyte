---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_xero Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceXero DataSource
---

# airbyte_source_xero (Data Source)

SourceXero DataSource

## Example Usage

```terraform
data "airbyte_source_xero" "my_source_xero" {
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

- `authentication` (Attributes) (see [below for nested schema](#nestedatt--configuration--authentication))
- `source_type` (String) must be one of ["xero"]
- `start_date` (String) UTC date and time in the format YYYY-MM-DDTHH:mm:ssZ. Any data with created_at before this data will not be synced.
- `tenant_id` (String) Enter your Xero organization's Tenant ID

<a id="nestedatt--configuration--authentication"></a>
### Nested Schema for `configuration.authentication`

Read-Only:

- `access_token` (String) Enter your Xero application's access token
- `client_id` (String) Enter your Xero application's Client ID
- `client_secret` (String) Enter your Xero application's Client Secret
- `refresh_token` (String) Enter your Xero application's refresh token
- `token_expiry_date` (String) The date-time when the access token should be refreshed


