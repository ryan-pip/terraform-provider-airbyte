---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_typesense Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationTypesense Resource
---

# airbyte_destination_typesense (Resource)

DestinationTypesense Resource

## Example Usage

```terraform
resource "airbyte_destination_typesense" "my_destination_typesense" {
  configuration = {
    api_key          = "...my_api_key..."
    batch_size       = "...my_batch_size..."
    destination_type = "typesense"
    host             = "...my_host..."
    port             = "...my_port..."
    protocol         = "...my_protocol..."
  }
  name         = "Wm Hane"
  workspace_id = "1d6c645b-08b6-4189-9baa-0fe1ade008e6"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `api_key` (String) Typesense API Key
- `destination_type` (String) must be one of ["typesense"]
- `host` (String) Hostname of the Typesense instance without protocol.

Optional:

- `batch_size` (String) How many documents should be imported together. Default 1000
- `port` (String) Port of the Typesense instance. Ex: 8108, 80, 443. Default is 443
- `protocol` (String) Protocol of the Typesense instance. Ex: http or https. Default is https


