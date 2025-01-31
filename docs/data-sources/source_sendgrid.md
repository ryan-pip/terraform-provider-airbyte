---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_sendgrid Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSendgrid DataSource
---

# airbyte_source_sendgrid (Data Source)

SourceSendgrid DataSource

## Example Usage

```terraform
data "airbyte_source_sendgrid" "my_source_sendgrid" {
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

- `apikey` (String) API Key, use <a href="https://app.sendgrid.com/settings/api_keys/">admin</a> to generate this key.
- `source_type` (String) must be one of ["sendgrid"]
- `start_time` (String) Start time in ISO8601 format. Any data before this time point will not be replicated.


