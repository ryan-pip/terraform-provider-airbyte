---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_google_webfonts Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGoogleWebfonts Resource
---

# airbyte_source_google_webfonts (Resource)

SourceGoogleWebfonts Resource

## Example Usage

```terraform
resource "airbyte_source_google_webfonts" "my_source_googlewebfonts" {
  configuration = {
    alt          = "...my_alt..."
    api_key      = "...my_api_key..."
    pretty_print = "...my_pretty_print..."
    sort         = "...my_sort..."
    source_type  = "google-webfonts"
  }
  name         = "Garrett Hoeger"
  secret_id    = "...my_secret_id..."
  workspace_id = "af90a26c-7cdc-4981-b068-981d6bb33cfa"
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

- `api_key` (String) API key is required to access google apis, For getting your's goto google console and generate api key for Webfonts
- `source_type` (String) must be one of ["google-webfonts"]

Optional:

- `alt` (String) Optional, Available params- json, media, proto
- `pretty_print` (String) Optional, boolean type
- `sort` (String) Optional, to find how to sort


