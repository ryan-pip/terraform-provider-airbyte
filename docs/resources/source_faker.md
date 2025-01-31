---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_faker Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceFaker Resource
---

# airbyte_source_faker (Resource)

SourceFaker Resource

## Example Usage

```terraform
resource "airbyte_source_faker" "my_source_faker" {
  configuration = {
    always_updated    = false
    count             = 4
    parallelism       = 1
    records_per_slice = 4
    seed              = 4
    source_type       = "faker"
  }
  name         = "Beth McKenzie"
  secret_id    = "...my_secret_id..."
  workspace_id = "1ec00221-b335-4d89-acb3-ecfda8d0c549"
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

- `count` (Number) How many users should be generated in total.  This setting does not apply to the purchases or products stream.
- `source_type` (String) must be one of ["faker"]

Optional:

- `always_updated` (Boolean) Should the updated_at values for every record be new each sync?  Setting this to false will case the source to stop emitting records after COUNT records have been emitted.
- `parallelism` (Number) How many parallel workers should we use to generate fake data?  Choose a value equal to the number of CPUs you will allocate to this source.
- `records_per_slice` (Number) How many fake records will be in each page (stream slice), before a state message is emitted?
- `seed` (Number) Manually control the faker random seed to return the same values on subsequent runs (leave -1 for random)


