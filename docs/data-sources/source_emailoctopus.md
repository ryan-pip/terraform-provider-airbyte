---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_emailoctopus Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceEmailoctopus DataSource
---

# airbyte_source_emailoctopus (Data Source)

SourceEmailoctopus DataSource

## Example Usage

```terraform
data "airbyte_source_emailoctopus" "my_source_emailoctopus" {
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

- `api_key` (String) EmailOctopus API Key. See the <a href="https://help.emailoctopus.com/article/165-how-to-create-and-delete-api-keys">docs</a> for information on how to generate this key.
- `source_type` (String) must be one of ["emailoctopus"]


