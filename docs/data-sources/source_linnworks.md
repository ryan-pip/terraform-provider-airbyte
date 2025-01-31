---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_linnworks Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceLinnworks DataSource
---

# airbyte_source_linnworks (Data Source)

SourceLinnworks DataSource

## Example Usage

```terraform
data "airbyte_source_linnworks" "my_source_linnworks" {
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

- `application_id` (String) Linnworks Application ID
- `application_secret` (String) Linnworks Application Secret
- `source_type` (String) must be one of ["linnworks"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
- `token` (String)


