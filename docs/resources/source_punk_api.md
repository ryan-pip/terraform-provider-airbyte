---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_punk_api Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePunkAPI Resource
---

# airbyte_source_punk_api (Resource)

SourcePunkAPI Resource

## Example Usage

```terraform
resource "airbyte_source_punk_api" "my_source_punkapi" {
  configuration = {
    brewed_after  = "MM-YYYY"
    brewed_before = "MM-YYYY"
    id            = 22
    source_type   = "punk-api"
  }
  name         = "Santos Langosh"
  secret_id    = "...my_secret_id..."
  workspace_id = "51bd76ee-eb51-48c4-9a1f-ad35512f06d4"
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

- `brewed_after` (String) To extract specific data with Unique ID
- `brewed_before` (String) To extract specific data with Unique ID
- `source_type` (String) must be one of ["punk-api"]

Optional:

- `id` (String) To extract specific data with Unique ID


