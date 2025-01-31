---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_vantage Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceVantage Resource
---

# airbyte_source_vantage (Resource)

SourceVantage Resource

## Example Usage

```terraform
resource "airbyte_source_vantage" "my_source_vantage" {
  configuration = {
    access_token = "...my_access_token..."
    source_type  = "vantage"
  }
  name         = "Veronica Pagac Sr."
  secret_id    = "...my_secret_id..."
  workspace_id = "38e1a735-ac26-4ae3-bbef-971a8f46bca1"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `access_token` (String) Your API Access token. See <a href="https://vantage.readme.io/reference/authentication">here</a>.
- `source_type` (String) must be one of ["vantage"]


