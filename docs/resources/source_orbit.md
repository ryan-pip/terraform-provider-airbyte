---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_orbit Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceOrbit Resource
---

# airbyte_source_orbit (Resource)

SourceOrbit Resource

## Example Usage

```terraform
resource "airbyte_source_orbit" "my_source_orbit" {
  configuration = {
    api_token   = "...my_api_token..."
    source_type = "orbit"
    start_date  = "...my_start_date..."
    workspace   = "...my_workspace..."
  }
  name         = "Dr. Nancy Ferry"
  secret_id    = "...my_secret_id..."
  workspace_id = "899f0c20-01e2-42cd-95cc-0584a184d76d"
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

- `api_token` (String) Authorizes you to work with Orbit workspaces associated with the token.
- `source_type` (String) must be one of ["orbit"]
- `workspace` (String) The unique name of the workspace that your API token is associated with.

Optional:

- `start_date` (String) Date in the format 2022-06-26. Only load members whose last activities are after this date.


