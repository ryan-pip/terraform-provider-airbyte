---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_qualaroo Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceQualaroo Resource
---

# airbyte_source_qualaroo (Resource)

SourceQualaroo Resource

## Example Usage

```terraform
resource "airbyte_source_qualaroo" "my_source_qualaroo" {
  configuration = {
    key         = "...my_key..."
    source_type = "qualaroo"
    start_date  = "2021-03-01T00:00:00.000Z"
    survey_ids = [
      "...",
    ]
    token = "...my_token..."
  }
  name         = "Grace Ankunding"
  secret_id    = "...my_secret_id..."
  workspace_id = "fbba5cce-ff5c-4b01-be51-e528a45ac82b"
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

- `key` (String) A Qualaroo token. See the <a href="https://help.qualaroo.com/hc/en-us/articles/201969438-The-REST-Reporting-API">docs</a> for instructions on how to generate it.
- `source_type` (String) must be one of ["qualaroo"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
- `token` (String) A Qualaroo token. See the <a href="https://help.qualaroo.com/hc/en-us/articles/201969438-The-REST-Reporting-API">docs</a> for instructions on how to generate it.

Optional:

- `survey_ids` (List of String) IDs of the surveys from which you'd like to replicate data. If left empty, data from all surveys to which you have access will be replicated.


