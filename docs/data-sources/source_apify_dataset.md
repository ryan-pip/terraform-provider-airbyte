---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_apify_dataset Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceApifyDataset DataSource
---

# airbyte_source_apify_dataset (Data Source)

SourceApifyDataset DataSource

## Example Usage

```terraform
data "airbyte_source_apify_dataset" "my_source_apifydataset" {
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

- `clean` (Boolean) If set to true, only clean items will be downloaded from the dataset. See description of what clean means in <a href="https://docs.apify.com/api/v2#/reference/datasets/item-collection/get-items">Apify API docs</a>. If not sure, set clean to false.
- `dataset_id` (String) ID of the dataset you would like to load to Airbyte.
- `source_type` (String) must be one of ["apify-dataset"]


